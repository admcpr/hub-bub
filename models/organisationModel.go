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
	SelectedRepo string
	Repositories []structs.Repository
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

	// m.RepositoryTable, cmd = m.RepositoryTable.Update(msg)
	m.repoList, cmd = m.repoList.Update(msg)

	return m, cmd
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

// View implements tea.Model
func (m OrganisationModel) View() string {
	var repoList = utils.BaseStyle.Render(m.repoList.View())

	var views = []string{repoList}

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
