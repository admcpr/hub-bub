package models

import (
	"hub-bub/structs"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RepoModel struct {
	repoSettingsTabs []structs.RepositorySettingsTab

	settingsTable table.Model

	help help.Model
	keys repoKeyMap

	activeTab int
	loaded    bool
	width     int
	height    int
}

func NewRepoModel(width, height int) RepoModel {
	return RepoModel{
		repoSettingsTabs: []structs.RepositorySettingsTab{},
		width:            width,
		height:           height,
		help:             help.New(),
		keys:             NewRepoKeyMap(),
	}
}

func (m RepoModel) Init() tea.Cmd {
	return nil
}

func (m *RepoModel) SelectRepo(RepositoryQuery structs.RepositoryQuery, width, height int) {
	m.repoSettingsTabs = structs.BuildRepoSettings(RepositoryQuery)

	m.width = width
	m.height = height
}

func (m *RepoModel) NextTab() {
	m.activeTab = min(m.activeTab+1, len(m.repoSettingsTabs)-1)
}

func (m *RepoModel) PreviousTab() {
	m.activeTab = max(m.activeTab-1, 0)
}

func (m RepoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
		}
		return m, nil
	}

	return m, cmd
}

func (m RepoModel) View() string {
	if m.repoSettingsTabs == nil || len(m.repoSettingsTabs) == 0 {
		return ""
	}

	m.buildSettingsTable()

	var tabs = m.RenderTabs()
	var settings = settingsStyle.Padding(0).Width(m.width - 2).Render(m.settingsTable.View())

	return lipgloss.JoinVertical(lipgloss.Left, tabs, settings)
}

func (m *RepoModel) buildSettingsTable() {
	var activeSettings = m.repoSettingsTabs[m.activeTab]

	columns := []table.Column{{Title: "", Width: 50}, {Title: "", Width: 11}}

	rows := make([]table.Row, len(activeSettings.Settings))
	for i, setting := range activeSettings.Settings {
		rows[i] = table.Row{setting.Name, setting.Value}
	}

	m.settingsTable = table.New(
		table.WithColumns(columns),
		table.WithRows(rows))
}

func (m RepoModel) RenderTabs() string {
	Tabs := []string{}
	for _, t := range m.repoSettingsTabs {
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

		if isLast {
			style = style.Border(border).Width((m.width / len(Tabs)) - 1)
		} else {
			style = style.Border(border).Width((m.width / len(Tabs)) - 3)
		}

		renderedTabs = append(renderedTabs, style.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	return row
}
