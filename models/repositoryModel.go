package models

import (
	"github.com/admcpr/hub-bub/structs"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RepositoryModel struct {
	repositorySettingsTabs []structs.RepositorySettingsTab

	settingsTable table.Model

	activeTab int
	loaded    bool
	width     int
	height    int
}

func NewRepositoryModel(width, height int) RepositoryModel {
	return RepositoryModel{
		repositorySettingsTabs: []structs.RepositorySettingsTab{},
		width:                  width,
		height:                 height,
	}
}

func (m RepositoryModel) Init() tea.Cmd {
	return nil
}

func (m *RepositoryModel) SelectRepo(RepositoryQuery structs.RepositoryQuery, width, height int) {
	m.repositorySettingsTabs = structs.BuildRepositorySettings(RepositoryQuery)

	m.width = width
	m.height = height
	m.buildSettingsTable(m.repositorySettingsTabs[m.activeTab])
}

func (m *RepositoryModel) NextTab() {
	m.activeTab = min(m.activeTab+1, len(m.repositorySettingsTabs)-1)
}

func (m *RepositoryModel) PreviousTab() {
	m.activeTab = max(m.activeTab-1, 0)
}

func (m RepositoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m RepositoryModel) View() string {
	var tabs = m.RenderTabs()
	var settings = settingsStyle.Padding(0).Width(m.width - 6).Render(m.settingsTable.View())

	return lipgloss.JoinVertical(lipgloss.Left, tabs, settings)
}

func (m *RepositoryModel) buildSettingsTable(tabSettings structs.RepositorySettingsTab) {

	columns := []table.Column{{Title: "", Width: 20}, {Title: "", Width: 11}}

	rows := make([]table.Row, len(tabSettings.Settings))
	for i, setting := range tabSettings.Settings {
		rows[i] = table.Row{setting.Name, setting.Value}
	}

	m.settingsTable = table.New(
		table.WithColumns(columns),
		table.WithRows(rows))
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
		// TODO: Calculate width of tabs correctly so they match m.width
		if isLast {
			style = style.Border(border).Width((m.width / len(Tabs)) - 4)
		} else {
			style = style.Border(border).Width((m.width / len(Tabs)) - 3)
		}

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
