package models

import tea "github.com/charmbracelet/bubbletea"

type YesNoModel struct {
	value bool
}

func (m YesNoModel) Init() tea.Cmd {
	return nil
}

func (m YesNoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m YesNoModel) View() string {
	if m.value {
		return "Yes"
	}
	return "No"
}
