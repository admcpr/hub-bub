package models

import (
	"log"

	"github.com/admcpr/hub-bub/messages"
	"github.com/admcpr/hub-bub/structs"
	"github.com/admcpr/hub-bub/utils"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
)

type OrganisationModel struct {
	Title        string
	Url          string
	Repositories []structs.Repository
	SelectedRepo RepositoryModel
	repoList     list.Model
	loaded       bool
	width        int
	height       int
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
		var columnStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder())
		var focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
		const divisor = 2

		m.height = msg.Height
		m.width = msg.Width

		if !m.loaded {
			columnStyle.Width(msg.Width / divisor)
			focusedStyle.Width(msg.Width / divisor)
			columnStyle.Height(msg.Height - divisor)
			focusedStyle.Height(msg.Height - divisor)
			m.initList()
			m.loaded = true
		}

	case messages.RepositoryListMsg:
		m.repoList = buildRepoListModel(msg.OrganizationQuery, m.width, m.height)
		m.Repositories = msg.OrganizationQuery.GetRepositories()
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!"),
			)
		case "esc":
			return MainModel[UserModelName], nil
		}
	}

	m.SelectedRepo.Update(msg)
	m.repoList, cmd = m.repoList.Update(msg)

	return m, cmd
}

// View implements tea.Model
func (m OrganisationModel) View() string {
	var repoList = utils.BaseStyle.Render(m.repoList.View())
	var repoTab = utils.BaseStyle.Render(m.SelectedRepo.View())

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
		items[i] = structs.NewListItem("repo.Node.Name"+repo.Node.Name, "repo.Node.Url")
	}

	list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Repositories"
	list.SetHeight(10)
	list.SetWidth(width)

	return list
}
