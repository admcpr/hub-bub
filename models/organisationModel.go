package models

import (
	"log"

	"hub-bub/messages"
	"hub-bub/structs"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
)

type OrganisationModel struct {
	Title     string
	Url       string
	RepoQuery structs.OrganizationQuery

	repoList  list.Model
	repoModel RepositoryModel

	loaded        bool
	width         int
	height        int
	tabsHaveFocus bool
}

func (m *OrganisationModel) panelWidth() int {
	return m.width / 2
}

func (m *OrganisationModel) getSelectedRepo() structs.RepositoryQuery {
	return m.RepoQuery.Organization.Repositories.Edges[m.repoList.Index()].Node
}

func (m *OrganisationModel) init(width, height int) {
	m.width = width
	m.height = height
	m.repoModel.width = m.panelWidth()
	m.repoModel.height = m.height
	m.loaded = true

	m.repoList = list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		m.panelWidth(),
		m.height,
	)

	m.repoModel = NewRepositoryModel(m.panelWidth(), m.height)
}

func (m OrganisationModel) Init() tea.Cmd {
	return nil
}

func (m OrganisationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		if !m.loaded {
			m.init(msg.Width, msg.Height)
		}
		return m, nil

	case messages.RepoListMsg:
		m.repoList = buildRepoListModel(msg.OrganizationQuery, m.width, m.height)
		m.RepoQuery = msg.OrganizationQuery
		m.repoModel.SelectRepo(m.getSelectedRepo(), m.panelWidth(), m.height)
		return m, nil

	case tea.KeyMsg:
		if m.tabsHaveFocus {
			// Pass messages to repository model
			switch msg.Type {
			case tea.KeyEsc:
				m.tabsHaveFocus = false
				return m, nil
			case tea.KeyRight:
				m.repoModel.NextTab()
				_, cmd = m.repoModel.Update(msg)
			case tea.KeyLeft:
				m.repoModel.PreviousTab()
				_, cmd = m.repoModel.Update(msg)
			}
		} else {
			switch msg.Type {
			case tea.KeyDown, tea.KeyUp:
				m.repoList, cmd = m.repoList.Update(msg)
				m.repoModel.SelectRepo(m.getSelectedRepo(), m.panelWidth(), m.height)
			case tea.KeyEnter:
				m.tabsHaveFocus = true
				m.repoModel.SelectRepo(m.getSelectedRepo(), m.panelWidth(), m.height)
			case tea.KeyEsc:
				return MainModel[UserModelName], nil
			}
		}
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, cmd
}

func (m OrganisationModel) View() string {
	var repoList = appStyle.Width(m.panelWidth() - 4).Render(m.repoList.View())
	var settings = appStyle.Width(m.panelWidth()).Render(m.repoModel.View())

	var views = []string{repoList, settings}

	return lipgloss.JoinHorizontal(lipgloss.Top, views...)
}

func (m OrganisationModel) GetRepositories() tea.Msg {
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

	list := list.New(items, list.NewDefaultDelegate(), width, height-titleHeight)
	list.Title = organizationQuery.Organization.Login
	list.SetStatusBarItemName("Repository", "Repositories")
	list.SetShowHelp(false)
	list.SetShowTitle(true)

	return list
}
