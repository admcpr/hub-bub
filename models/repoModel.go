package models

import (
	"hub-bub/keyMaps"
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
	keys keyMaps.RepoKeyMap

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
		keys:             keyMaps.NewRepoKeyMap(),
	}
}

func (m RepoModel) Init() tea.Cmd {
	return nil
}

func (m *RepoModel) SelectRepo(RepositoryQuery structs.RepositoryQuery, width, height int) {
	m.repoSettingsTabs = structs.BuildRepoSettings(RepositoryQuery)
	m.settingsTable = NewSettingsTable(m.repoSettingsTabs[m.activeTab], width)

	m.width = width
	m.height = height
}

func (m *RepoModel) NextTab() {
	m.activeTab = min(m.activeTab+1, len(m.repoSettingsTabs)-1)
	m.settingsTable = NewSettingsTable(m.repoSettingsTabs[m.activeTab], m.width)
}

func (m *RepoModel) PreviousTab() {
	m.activeTab = max(m.activeTab-1, 0)
	m.settingsTable = NewSettingsTable(m.repoSettingsTabs[m.activeTab], m.width)
}

func (m RepoModel) Update(msg tea.Msg) (RepoModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.loaded = true
		}
		return m, nil
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDown:
			m.settingsTable.MoveDown(1)
		case tea.KeyUp:
			m.settingsTable.MoveUp(1)
		}
	}

	return m, cmd
}

func (m RepoModel) View() string {
	if m.repoSettingsTabs == nil || len(m.repoSettingsTabs) == 0 {
		return ""
	}

	settingsStyle := appStyle.Copy().Border(settingsBorder()).
		BorderForeground(blueLighter).Padding(0).Margin(0)

	var tabs = RenderTabs(m.repoSettingsTabs, m.width, m.activeTab)
	var settings = settingsStyle.Width(m.width - 2).Render(m.settingsTable.View())
	// var settings = settingsStyle.Width(m.width - 2).Height(m.height - 7).Render(m.settingsTable.View())

	return lipgloss.JoinVertical(lipgloss.Left, tabs, settings)
}

func NewSettingsTable(activeSettings structs.RepositorySettingsTab, width int) table.Model {
	widthWithoutBorder := width - 2
	quarterWidth := quarter(widthWithoutBorder)

	columns := []table.Column{
		{Title: "Setting", Width: (widthWithoutBorder - quarterWidth)},
		{Title: "Value", Width: quarterWidth}}

	rows := make([]table.Row, len(activeSettings.Settings))
	for i, setting := range activeSettings.Settings {
		rows[i] = table.Row{setting.Name, setting.Value}
	}

	return table.New(table.WithColumns(columns),
		table.WithRows(rows), table.WithFocused(true), table.WithStyles(GetTableStyles()))
}

func GetTableStyles() table.Styles {
	return table.Styles{
		Selected: lipgloss.NewStyle().Bold(true).Foreground(pink),
		Header:   lipgloss.NewStyle().Bold(true).Foreground(blue).Underline(true).Padding(0, 0, 1, 0),
		Cell:     lipgloss.NewStyle().Padding(0),
	}
}
