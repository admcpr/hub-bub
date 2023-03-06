package models

import (
	"github.com/admcpr/hub-bub/messages"
	"github.com/admcpr/hub-bub/structs"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RepositoryModel struct {
	repositorySettingsTabs []structs.RepositorySettingsTab

	settingList list.Model

	activeTab     int
	loaded        bool
	width         int
	height        int
	tabsHaveFocus bool
}

func NewRepositoryModel(width, height int) RepositoryModel {
	return RepositoryModel{
		repositorySettingsTabs: []structs.RepositorySettingsTab{},
		settingList:            list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		width:                  width,
		height:                 height,
	}
}

func (m RepositoryModel) Init() tea.Cmd {
	return nil
}

func (m RepositoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		if !m.loaded {
			m.loaded = true
		}
		return m, nil

	case messages.RepoSelectedMsg:
		m.repositorySettingsTabs = structs.BuildRepositorySettings(msg.RepositoryQuery)
		return m, nil

	// case messages.RepoListMsg:
	// 	// m.repositorySettingsTabs = structs.BuildRepositorySettings(m.RepoQuery.Organization.Repositories.Edges[m.repoList.Index()].Node)
	// 	return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			m.tabsHaveFocus = false
			return m, nil
		case tea.KeyRight:
			m.activeTab = min(m.activeTab+1, len(m.repositorySettingsTabs)-1)
		case tea.KeyLeft:
			m.activeTab = max(m.activeTab-1, 0)
		}
	}

	m.buildSettingListModel(m.repositorySettingsTabs[m.activeTab], m.width, m.height)

	return m, cmd
}

func (m RepositoryModel) View() string {
	var tabs = m.RenderTabs()
	var settings = settingsStyle.Render(m.settingList.View())

	return lipgloss.JoinVertical(lipgloss.Left, tabs, settings)
}

func (m *RepositoryModel) buildSettingListModel(tabSettings structs.RepositorySettingsTab, width, height int) {
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

func (m RepositoryModel) RenderTabs() string {
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
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	return row
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
