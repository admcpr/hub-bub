package models

import (
	"fmt"
	"io"
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
	Title     string
	Url       string
	RepoQuery structs.OrganizationQuery

	repositorySettingsTabs []structs.RepositorySettingsTab

	repoList    list.Model
	settingList list.Model

	activeTab     int
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
	m.settingList = list.New(
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
		m.repositorySettingsTabs = structs.BuildRepositorySettings(m.RepoQuery.Organization.Repositories.Edges[m.repoList.Index()].Node)
		m.settingList = buildSettingListModel(m.repositorySettingsTabs[m.activeTab], m.width, m.height)
		return m, nil

	case tea.KeyMsg:
		if m.tabsHaveFocus {
			// Pass messages to repository model
			switch msg.Type {
			case tea.KeyEsc:
				m.tabsHaveFocus = false
				return m, nil
			case tea.KeyRight:
				m.activeTab = min(m.activeTab+1, len(m.repositorySettingsTabs)-1)
				m.settingList = buildSettingListModel(m.repositorySettingsTabs[m.activeTab], m.width, m.height)
			case tea.KeyLeft:
				m.activeTab = max(m.activeTab-1, 0)
				m.settingList = buildSettingListModel(m.repositorySettingsTabs[m.activeTab], m.width, m.height)
			}
		} else {
			switch msg.Type {
			case tea.KeyDown, tea.KeyUp:
				m.repositorySettingsTabs = structs.BuildRepositorySettings(m.RepoQuery.Organization.Repositories.Edges[m.repoList.Index()].Node)
				m.settingList = buildSettingListModel(m.repositorySettingsTabs[m.activeTab], m.width, m.height)
			case tea.KeyEnter:
				m.tabsHaveFocus = true
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
	var repoList = appStyle.Render(m.repoList.View())
	var settingList = lipgloss.JoinVertical(lipgloss.Left, m.Tabs(), appStyle.Render(m.settingList.View()))

	var views = []string{repoList, settingList}

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

	list := list.New(items, list.NewDefaultDelegate(), width, height-titleHeight)
	list.Title = "Repositories"
	list.SetShowHelp(false)
	list.SetShowTitle(true)

	return list
}

func buildSettingListModel(tabSettings structs.RepositorySettingsTab, width, height int) list.Model {
	items := make([]list.Item, len(tabSettings.Settings))
	for i, setting := range tabSettings.Settings {
		items[i] = structs.NewListItem(setting.Name, setting.Value)
	}

	list := list.New(items, itemDelegate{}, width, height-titleHeight-4)
	list.Title = tabSettings.Name
	list.SetShowTitle(false)
	list.SetShowStatusBar(false)

	return list
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 1 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(structs.ListItem)
	if !ok {
		return
	}

	// str := fmt.Sprintf("%s > %s", i.Title(), i.Description())

	statusNugget := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Padding(0, 1)

	// statusBarStyle := lipgloss.NewStyle().
	// 	Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
	// 	Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	// statusStyle := lipgloss.NewStyle().
	// 	Inherit(statusBarStyle).
	// 	Foreground(lipgloss.Color("#FFFDF5")).
	// 	Background(lipgloss.Color("#FF5F87")).
	// 	Padding(0, 1).
	// 	MarginRight(1)

	encodingStyle := statusNugget.Copy().
		Background(lipgloss.Color(pink)).
		Align(lipgloss.Right)

	// str := lipgloss.JoinHorizontal(lipgloss.Left, i.Title(), i.Description())

	title := statusNugget.Render(i.Title())
	description := encodingStyle.Render(i.Description())

	// fn := itemStyle.Render
	// if index == m.Index() {
	// 	fn = func(s string) string {
	// 		return selectedItemStyle.Render("> " + s)
	// 	}
	// }
	fn := lipgloss.JoinHorizontal

	fmt.Fprint(w, fn(lipgloss.Left, title, description))
}

func (m OrganisationModel) Tabs() string {
	Tabs := []string{"Overview", "Features", "PRs & Default Branch", "Security", "Wiki", "Settings"}

	var renderedTabs []string

	for i, t := range Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(Tabs)-1, i == m.activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	return row
}
