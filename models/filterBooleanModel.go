package models

import (
	"hub-bub/structs"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterBooleanModel struct {
	Title string
	input textinput.Model
}

func NewFilterBooleanModel(title string, value bool) FilterBooleanModel {
	m := FilterBooleanModel{
		Title: title,
		input: textinput.New(),
	}

	m.input.SetValue(structs.YesNo(value))
	m.input.Focus()

	return m
}

func (m FilterBooleanModel) Init() tea.Cmd {
	return m.Focus()
}

func (m FilterBooleanModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m FilterBooleanModel) View() string {
	return m.Title + " " + m.input.View()
}

func (m *FilterBooleanModel) GetValue() bool {
	return m.input.Value() == "Yes"
}

func (m *FilterBooleanModel) Focus() tea.Cmd {
	return m.input.Focus()
}
