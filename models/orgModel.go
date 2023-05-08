package models

import (
	"log"

	"hub-bub/keyMaps"
	"hub-bub/messages"
	"hub-bub/structs"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
)

type OrgModel struct {
	Title     string
	RepoQuery structs.OrganizationQuery

	repoList  list.Model
	repoModel RepoModel
	help      help.Model
	keys      keyMaps.OrgKeyMap

	loaded        bool
	width         int
	height        int
	tabsHaveFocus bool
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
	}
}

func (m *OrgModel) helpView() string {
	if m.tabsHaveFocus {
		return m.repoModel.help.View(m.repoModel.keys)
	}

	return m.help.View(m.keys)
}

func (m *OrgModel) getSelectedRepo() structs.RepositoryQuery {
	return m.RepoQuery.Organization.Repositories.Edges[m.repoList.Index()].Node
}

func (m *OrgModel) init(width, height int) {
	m.loaded = true
}

func (m OrgModel) Init() tea.Cmd {
	return nil
}

func (m OrgModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		if !m.loaded {
			m.init(msg.Width, msg.Height)
		}
		return m, nil

	case messages.RepoListMsg:
		m.repoList = buildRepoListModel(msg.OrganizationQuery, m.width, m.height)
		m.RepoQuery = msg.OrganizationQuery
		m.repoModel.SelectRepo(m.getSelectedRepo(), half(m.width), m.height)
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
			}
		} else {
			switch msg.String() {
			case tea.KeyEnter.String():
				m.tabsHaveFocus = true
			case tea.KeyEsc.String():
				if !m.repoList.FilteringEnabled() {
					return MainModel[UserModelName], nil
				}
			case "ctrl+c", "q":
				if !m.repoList.FilteringEnabled() {
					return m, tea.Quit
				}
			case tea.KeyUp.String(), tea.KeyDown.String(), tea.KeyLeft.String(), tea.KeyRight.String():
				m.repoModel.SelectRepo(m.getSelectedRepo(), half(m.width), m.height)
			}
		}
	}

	if m.tabsHaveFocus {
		_, cmd = m.repoModel.Update(msg)
	} else {
		m.repoList, cmd = m.repoList.Update(msg)
	}

	return m, cmd
}

func (m OrgModel) View() string {
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
		"first": graphql.Int(30),
	}
	err = client.Query("OrganizationRepositories", &organizationQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	return messages.RepoListMsg{OrganizationQuery: organizationQuery}
}

func buildRepoListModel(organizationQuery structs.OrganizationQuery, width, height int) list.Model {
	edges := organizationQuery.Organization.Repositories.Edges
	items := make([]list.Item, len(edges))
	for i, repo := range edges {
		items[i] = structs.NewListItem(repo.Node.Name, repo.Node.Url)
	}

	list := list.New(items, list.NewDefaultDelegate(), width, height-2)
	list.Title = organizationQuery.Organization.Login
	list.SetStatusBarItemName("Repository", "Repositories")
	list.SetShowHelp(false)
	list.SetShowTitle(true)

	return list
}
