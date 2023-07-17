package models

import (
	"hub-bub/messages"
	"hub-bub/structs"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterBoolModel struct {
	Tab   string
	Title string
	input textinput.Model
}

func NewFilterBoolModel(tab, title string, value bool) FilterBoolModel {
	m := FilterBoolModel{
		Tab:   tab,
		Title: title,
		input: textinput.New(),
	}

	m.input.SetValue(structs.YesNo(value))
	m.input.Focus()

	return m
}

func (m FilterBoolModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m FilterBoolModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEnter.String():
			return m, m.Confirm
		case tea.KeyEsc.String():
			return m, m.Cancel
		case "y", "Y":
			m.input.SetValue("Yes")
		case "n", "N":
			m.input.SetValue("No")
		}
	}

	return m, cmd
}

func (m FilterBoolModel) View() string {
	return m.Title + " " + m.input.View()
}

func (m *FilterBoolModel) GetValue() bool {
	return m.input.Value() == "Yes"
}

func (m *FilterBoolModel) Focus() tea.Cmd {
	return m.input.Focus()
}

func (m FilterBoolModel) Cancel() tea.Msg {
	return messages.NewCancelFilterMsg(structs.NewFilterBool(m.Tab, m.Title, false))
}

func (m FilterBoolModel) Confirm() tea.Msg {
	return messages.NewAddFilterMsg(structs.NewFilterBool(m.Tab, m.Title, m.GetValue()))
}
