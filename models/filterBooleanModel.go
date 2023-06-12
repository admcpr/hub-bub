package models

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterBooleanModel struct {
	Title string
	input textinput.Model
}

func (m FilterBooleanModel) Init() tea.Cmd {
	return nil
}

func (m FilterBooleanModel) Update(msg tea.Msg) (FilterBooleanModel, tea.Cmd) {
	return m, nil
}

func (m FilterBooleanModel) View() string {
	return m.Title + "\n\n" + m.input.View()
}
