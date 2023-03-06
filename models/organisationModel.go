package models

import (
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
	repoModel   RepositoryModel

	loaded        bool
	width         int
	height        int
	tabsHaveFocus bool
}

func (m *OrganisationModel) panelWidth() int {
	return m.width / 2
}

func (m *OrganisationModel) init() {
	m.repoList = list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		// m.width,
		// m.height,
		0,
		0,
	)
	m.settingList = list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		// m.width,
		// m.height,
		0,
		0,
	)
	m.repoModel = NewRepositoryModel(m.panelWidth(), m.height)
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
		m.repoModel.SelectRepo(m.getSelectedRepo())
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
			case tea.KeyLeft:
				m.repoModel.PreviousTab()
			}
		} else {
			switch msg.Type {
			case tea.KeyDown, tea.KeyUp:
				m.repoModel.SelectRepo(m.getSelectedRepo())
			case tea.KeyEnter:
				m.tabsHaveFocus = true
				m.repoModel.SelectRepo(m.getSelectedRepo())
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

	// m.buildSettingListModel(m.repositorySettingsTabs[m.activeTab], m.width, m.height)

	return m, cmd
}

// View implements tea.Model
func (m OrganisationModel) View() string {
	var repoList = appStyle.Width((m.width / 2) - 4).Render(m.repoList.View())
	var settingList = lipgloss.JoinVertical(lipgloss.Left, m.Tabs(), settingsStyle.Width(m.width/2).Render(m.settingList.View()))

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
	list.Title = organizationQuery.Organization.Login
	list.SetStatusBarItemName("Repository", "Repositories")
	list.SetShowHelp(false)
	list.SetShowTitle(true)

	return list
}

func (m *OrganisationModel) buildSettingListModel(tabSettings structs.RepositorySettingsTab, width, height int) {
	items := make([]list.Item, len(tabSettings.Settings))
	for i, setting := range tabSettings.Settings {
		items[i] = structs.NewListItem(setting.Name, setting.Value)
	}

	m.settingList = list.New(items, itemDelegate{}, width, height-titleHeight-4)
	m.settingList.Title = tabSettings.Name
	m.settingList.SetShowHelp(false)
	m.settingList.SetShowTitle(false)
	m.settingList.SetShowStatusBar(false)
}

func (m OrganisationModel) Tabs() string {
	Tabs := []string{}
	for _, t := range m.repositorySettingsTabs {
		Tabs = append(Tabs, t.Name)
	}

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
		// TODO: Calculate width of tabs correctly so they match m.width
		style = style.Border(border).Width((m.width / 2 / len(Tabs)) - 1)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	return row
}
