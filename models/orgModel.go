package models

import (
	"fmt"
	"log"

	"hub-bub/keyMaps"
	"hub-bub/messages"
	"hub-bub/structs"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
)

type OrgModel struct {
	Title        string
	Filters      []structs.RepositoryFilter
	Repositories []structs.Repository

	repoList  list.Model
	repoModel RepoModel
	help      help.Model
	keys      keyMaps.OrgKeyMap

	width         int
	height        int
	loaded        bool
	tabsHaveFocus bool
	getting       bool
	spinner       spinner.Model
}

func NewOrgModel(title string, width, height int) OrgModel {
	return OrgModel{
		Title:     title,
		width:     width,
		height:    height,
		help:      help.New(),
		keys:      keyMaps.NewOrgKeyMap(),
		repoModel: NewRepoModel(width/2, height),
		repoList:  list.New([]list.Item{}, list.NewDefaultDelegate(), width/2, height),
		Filters:   []structs.RepositoryFilter{},
		getting:   true,
		spinner:   spinner.New(spinner.WithSpinner(spinner.Pulse)),
	}
}

func (m *OrgModel) FilteredRepositories() []structs.Repository {
	if len(m.Filters) == 0 {
		return m.Repositories
	}
	filteredRepos := []structs.Repository{}
	// TODO: This is gonna get slow, fast, for big orgs. Faster pls.
	for _, repo := range m.Repositories {
		for _, filter := range m.Filters {
			for _, tab := range repo.SettingsTabs {
				if tab.Name == filter.Tab {
					for _, setting := range tab.Settings {
						if setting.GetName() == filter.SettingGetter.GetName() && setting.GetValue() == filter.SettingGetter.GetValue() {
							filteredRepos = append(filteredRepos, repo)
						}
					}
				}
			}
		}
	}
	return filteredRepos
}

func (m *OrgModel) UpdateRepositories(oq structs.OrganizationQuery) {
	edges := oq.Organization.Repositories.Edges
	m.Repositories = make([]structs.Repository, len(edges))
	items := make([]list.Item, len(edges))
	for i, repoQuery := range edges {
		repo := structs.NewRepository(repoQuery.Node)
		m.Repositories[i] = repo
		items[i] = structs.NewListItem(repo.Name, repo.Url)
	}

	m.UpdateRepoList()
	m.getting = false
}

func (m *OrgModel) UpdateRepoList() {
	filteredRepositories := m.FilteredRepositories()
	items := make([]list.Item, len(filteredRepositories))
	for i, repo := range m.FilteredRepositories() {
		items[i] = structs.NewListItem(repo.Name, repo.Url)
	}

	list := list.New(items, defaultDelegate, m.width, m.height-2)
	list.Title = "Organization: " + m.Title
	list.Styles.Title = titleStyle
	list.SetStatusBarItemName("Repository", "Repositories")
	list.SetShowHelp(false)
	list.SetShowTitle(true)

	m.repoList = list

	m.repoModel.SelectRepo(m.Repositories[0], half(m.width), m.height)
}

func (m *OrgModel) helpView() string {
	if m.tabsHaveFocus {
		return m.repoModel.help.View(m.repoModel.keys)
	}

	return m.help.View(m.keys)
}

func (m OrgModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m OrgModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
		}
		return m, nil

	case messages.RepoListMsg:
		m.UpdateRepositories(msg.OrganizationQuery)
		return m, nil

	case tea.KeyMsg:
		if m.tabsHaveFocus {
			switch msg.Type {
			case tea.KeyEsc:
				m.tabsHaveFocus = false
				return m, nil
			case tea.KeyRight:
				m.repoModel.NextTab()
			case tea.KeyLeft:
				m.repoModel.PreviousTab()
			case tea.KeyEnter:
				m.repoModel.ToggleFilterEditor()
			}
			m.repoModel, cmd = m.repoModel.Update(msg)
		} else {
			switch msg.String() {
			case tea.KeyEnter.String():
				m.tabsHaveFocus = true
			case tea.KeyEsc.String():
				if !m.repoList.SettingFilter() {
					return MainModel[UserModelName], nil
				}
				m.repoList, cmd = m.repoList.Update(msg)
			case "ctrl+c", "q":
				if !m.repoList.SettingFilter() {
					return m, tea.Quit
				}
			case tea.KeyUp.String(), tea.KeyDown.String():
				m.repoList, cmd = m.repoList.Update(msg)
				m.repoModel.SelectRepo(m.Repositories[m.repoList.Index()], half(m.width), m.height)
			default:
				m.repoList, cmd = m.repoList.Update(msg)
			}
		}

	default:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, cmd
}

func (m OrgModel) View() string {
	if m.getting {
		return fmt.Sprintf("%s getting repos ...", m.spinner.View())
	}

	var repoList = appStyle.Width(half(m.width)).Render(m.repoList.View())
	var settings = appStyle.Width(half(m.width)).Render(m.repoModel.View())
	help := m.helpView()
	var rightPanel = lipgloss.JoinVertical(lipgloss.Center, settings, help)

	var views = []string{repoList, rightPanel}

	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}

func (m OrgModel) GetRepositories() tea.Msg {
	client, err := gh.GQLClient(nil)
	if err != nil {
		return messages.AuthenticationErrorMsg{Err: err}
	}

	var organizationQuery = structs.OrganizationQuery{}

	variables := map[string]interface{}{
		"login": graphql.String(m.Title),
		"first": graphql.Int(100),
	}
	err = client.Query("OrganizationRepositories", &organizationQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	return messages.RepoListMsg{OrganizationQuery: organizationQuery}
}
