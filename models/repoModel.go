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

	settingsStyle := appStyle.Copy().Border(settingsBorder()).
		BorderForeground(lipgloss.Color(pink)).Padding(0).Margin(0)

	m.buildSettingsTable()

	var tabs = RenderTabs(m.repoSettingsTabs, m.width, m.activeTab)
	var settings = settingsStyle.Padding(0).Width(m.width - 2).Height(m.height - 7).Render(m.settingsTable.View())

	return lipgloss.JoinVertical(lipgloss.Left, tabs, settings)
}

func (m *RepoModel) buildSettingsTable() {
	activeSettings := m.repoSettingsTabs[m.activeTab]
	widthWithoutBorder := m.width - 2
	quarterWidth := quarter(widthWithoutBorder)

	columns := []table.Column{{Title: "", Width: (widthWithoutBorder - quarterWidth)}, {Title: "", Width: quarterWidth}}

	rows := make([]table.Row, len(activeSettings.Settings))
	for i, setting := range activeSettings.Settings {
		rows[i] = table.Row{setting.Name, setting.Value}
	}

	m.settingsTable = table.New(
		table.WithColumns(columns),
		table.WithRows(rows))
}
