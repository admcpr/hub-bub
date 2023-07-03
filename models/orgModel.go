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

type Focus int

const (
	focusList Focus = iota
	focusTabs
	focusFilter
)

func (f Focus) Next() Focus {
	switch f {
	case focusList:
		return focusTabs
	default:
		return focusFilter
	}
}

func (f Focus) Prev() Focus {
	switch f {
	case focusFilter:
		return focusTabs
	default:
		return focusList
	}
}

type OrgModel struct {
	Title   string
	Filters []structs.RepositoryFilter

	Repositories []structs.Repository

	repoList  list.Model
	repoModel tea.Model
	help      help.Model
	keys      keyMaps.OrgKeyMap

	focus   Focus
	width   int
	height  int
	loaded  bool
	getting bool
	spinner spinner.Model
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
		focus:     focusList,
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
						if setting.Name == filter.Setting.Name && setting.String() == filter.Setting.String() {
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

	// m.repoModel.SelectRepo(m.Repositories[0], half(m.width), m.height)
}

func (m *OrgModel) helpView() string {
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
		// Handle navigation Enter, Esc to select, deselect
		switch msg.Key {
		case tea.KeyEnter:
			//Enter selects so we can navigate to the next focus
			m.focus = m.focus.Next()
		case tea.KeyEsc:
			// Esc goes back so we can navigate to the previous focus, or go to the previous model
			if m.focus == focusList {
				if !m.repoList.SettingFilter() {
					return MainModel[UserModelName], nil
				}
			}
			m.focus = m.focus.Prev()
		}
	}

	// case tea.KeyMsg:
	// 	switch m.focus {
	// 	case focusList:
	// 		var tabCmd tea.Cmd
	// 		m, cmd = m.UpdateList(msg)

	// 		m, tabCmd = UpdateTabs(&m, msg)
	// 		return m, tea.Batch(cmd, tabCmd)
	// 	case focusTabs:
	// 		m, cmd = m.UpdateTabs(msg)
	// 	case focusFilter:
	// 		m, cmd = m.UpdateFilter(msg)
	// 	}
	// }

	if m.getting {
		m.spinner, cmd = m.spinner.Update(msg)
	}

	return m, cmd
}

func (m OrgModel) UpdateList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case tea.KeyEnter.String():
		m.focus = focusTabs
	case tea.KeyEsc.String():

		m.repoList, cmd = m.repoList.Update(msg)
	case "ctrl+c", "q":
		if !m.repoList.SettingFilter() {
			return m, tea.Quit
		}
	default:
		m.repoList, cmd = m.repoList.Update(msg)
	}

	return m, cmd
}

func (m OrgModel) UpdateTabs(msg tea.KeyMsg) (OrgModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg.Type {
	case tea.KeyEsc:
		m.focus = focusList
		return m, nil
	case tea.KeyEnter:
		// m.repoModel.ToggleFilterEditor()
	case tea.KeyUp, tea.KeyDown:
		// m.repoModel.SelectRepo(m.Repositories[m.repoList.Index()], half(m.width), m.height)
	}
	m.repoModel, cmd = m.repoModel.Update(msg)
	return m, cmd
}

func (m OrgModel) UpdateFilter(msg tea.KeyMsg) (OrgModel, tea.Cmd) {
	var cmd tea.Cmd

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
