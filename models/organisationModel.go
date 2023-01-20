package models

import (
	"fmt"
	"log"

	"github.com/admcpr/hub-bub/messages"
	"github.com/admcpr/hub-bub/structs"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
)

type OrganisationModel struct {
	Title         string
	Url           string
	RepoQuery     structs.OrganizationQuery
	SelectedRepo  RepositoryModel
	repoList      list.Model
	loaded        bool
	width         int
	height        int
	tabsHaveFocus bool
}

func (m *OrganisationModel) initList() {
	m.repoList = list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		m.width,
		m.height,
	)
}

func (m OrganisationModel) Init() tea.Cmd {
	return nil
}

func (m OrganisationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		if !m.loaded {
			m.initList()
			m.loaded = true
		}
		return m, nil

	case messages.RepositoryListMsg:
		m.repoList = buildRepoListModel(msg.OrganizationQuery, m.width, m.height)
		m.RepoQuery = msg.OrganizationQuery
		m.SelectedRepo = NewRepositoryModel(m.RepoQuery.Organization.Repositories.Edges[m.repoList.Index()].Node, m.width/2)
		return m, nil

	case tea.KeyMsg:
		if m.tabsHaveFocus {
			// Pass messages to repository model
			switch msg.Type {
			case tea.KeyEsc:
				m.tabsHaveFocus = false
				return m, nil
			}
			_, cmd = m.SelectedRepo.Update(msg)
		} else {
			switch msg.Type {
			case tea.KeyDown, tea.KeyUp:
				m.SelectedRepo = NewRepositoryModel(m.RepoQuery.Organization.Repositories.Edges[m.repoList.Index()].Node, m.width/2)
			case tea.KeyEnter:
				m.tabsHaveFocus = true
				return m, nil
			case tea.KeyEsc:
				return MainModel[UserModelName], nil
			}
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			}
			m.repoList, cmd = m.repoList.Update(msg)
		}
	}

	return m, cmd
}

// View implements tea.Model
func (m OrganisationModel) View() string {
	var docStyle = lipgloss.NewStyle().Margin(1, 2).Width(m.width / 2).Height(m.height)

	var repoList = docStyle.Render(m.repoList.View())
	var repoTab = docStyle.Render(m.SelectedRepo.View())

	var views = []string{repoList, repoTab}

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

	return messages.RepositoryListMsg{OrganizationQuery: organizationQuery}
}

func buildRepoListModel(organizationQuery structs.OrganizationQuery, width, height int) list.Model {
	edges := organizationQuery.Organization.Repositories.Edges
	items := make([]list.Item, len(edges))
	for i, repo := range edges {
		items[i] = structs.NewListItem(repo.Node.Name, repo.Node.Url)
	}

	list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	list.Title = fmt.Sprintf("Repositories h:%d w:%d", height, width)
	list.SetHeight(height)
	list.SetWidth(width)

	return list
}
