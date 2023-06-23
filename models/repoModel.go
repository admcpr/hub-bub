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
	repository  structs.Repository
	filterModel FilterModel // TODO:

	settingsTable table.Model

	help help.Model
	keys keyMaps.RepoKeyMap

	activeTab        int
	showFilterEditor bool
	loaded           bool
	hasFocus         bool
	width            int
	height           int
}

func NewRepoModel(width, height int) RepoModel {
	return RepoModel{
		repository: structs.Repository{},
		width:      width,
		height:     height,
		help:       help.New(),
		keys:       keyMaps.NewRepoKeyMap(),
	}
}

func (m RepoModel) HasFocus() bool {
	return m.hasFocus
}

func (m *RepoModel) ToggleFocus() {
	m.hasFocus = !m.hasFocus
}

func (m RepoModel) Init() tea.Cmd {
	return nil
}

func (m *RepoModel) SelectRepo(repository structs.Repository, width, height int) {
	m.repository = repository
	m.settingsTable = NewSettingsTable(m.repository.SettingsTabs[m.activeTab].Settings, width)

	m.width = width
	m.height = height
}

func (m *RepoModel) SelectTab(index int) {
	m.activeTab = index
	m.settingsTable = NewSettingsTable(m.repository.SettingsTabs[m.activeTab].Settings, m.width)
}

func (m *RepoModel) ToggleFilterEditor() {
	if !m.showFilterEditor {
		tab := m.repository.SettingsTabs[m.activeTab]
		index := m.settingsTable.Cursor()

		m.filterModel = NewFilterModel(tab.Name, tab.Settings[index])
	}

	m.showFilterEditor = !m.showFilterEditor
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
		case tea.KeyRight:
			m.SelectTab(min(m.activeTab+1, len(m.repository.SettingsTabs)-1))
		case tea.KeyLeft:
			m.SelectTab(max(m.activeTab-1, 0))
		case tea.KeyDown:
			m.settingsTable.MoveDown(1)
		case tea.KeyUp:
			m.settingsTable.MoveUp(1)
		case tea.KeyEsc:
			m.showFilterEditor = false
		}
	}

	m.filterModel, cmd = m.filterModel.Update(msg)

	return m, cmd
}

func (m RepoModel) View() string {
	if m.repository.SettingsTabs == nil || len(m.repository.SettingsTabs) == 0 {
		// Can this ever happen ????
		return ""
	}

	settingsStyle := appStyle.Copy().Border(settingsBorder()).
		BorderForeground(blueLighter).Padding(0).Margin(0)

	var tabs = RenderTabs(m.repository.SettingsTabs, m.width, m.activeTab)
	if m.showFilterEditor {
		filter := lipgloss.NewStyle().Width(m.width - 2).Height(m.height - 7).Render(m.filterModel.View())
		return lipgloss.JoinVertical(lipgloss.Left, tabs, filter)
	} else {
		settings := settingsStyle.Width(m.width - 2).Height(m.height - 7).Render(m.settingsTable.View())
		return lipgloss.JoinVertical(lipgloss.Left, tabs, settings)
	}
}

func NewSettingsTable(activeSettings []structs.SettingGetter, width int) table.Model {
	widthWithoutBorder := width - 2
	quarterWidth := quarter(widthWithoutBorder)

	columns := []table.Column{
		{Title: "Setting", Width: (widthWithoutBorder - quarterWidth)},
		{Title: "Value", Width: quarterWidth}}

	rows := make([]table.Row, len(activeSettings))
	for i, setting := range activeSettings {
		rows[i] = table.Row{setting.GetName(), setting.GetValue()}
	}

	return table.New(table.WithColumns(columns),
		table.WithRows(rows), table.WithFocused(true), table.WithStyles(GetTableStyles()))
}

func GetTableStyles() table.Styles {
	return table.Styles{
		Selected: lipgloss.NewStyle().Bold(true).Background(pink),
		Header: lipgloss.NewStyle().Bold(true).Foreground(blue).BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).BorderForeground(blueLighter),
		Cell: lipgloss.NewStyle().Padding(0),
	}
}
