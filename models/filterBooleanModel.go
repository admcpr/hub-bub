package models

import (
	"hub-bub/messages"
	"hub-bub/structs"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterBool struct {
	Tab   string
	Name  string
	Value bool
}

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
	return textinput.Blink
}

func (m FilterBooleanModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEnter.String():
			m.input.Blur()
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

func (m FilterBooleanModel) View() string {
	return m.Title + " " + m.input.View()
}

func (m *FilterBooleanModel) GetValue() bool {
	return m.input.Value() == "Yes"
}

func (m *FilterBooleanModel) Focus() tea.Cmd {
	return m.input.Focus()
}

func (m FilterBooleanModel) Cancel() tea.Msg {
	return messages.FilterCancelMsg{FilterName: m.Title}
}

func (m FilterBooleanModel) Confirm() tea.Msg {
	filter := structs.Filter{Title: m.Title, Value: m.GetValue()}
	msg := messages.FilterMsg{FilterName: m.Title, Action: structs.ConfirmAction}
}

func (m FilterBooleanModel) SendMsg(action string) tea.Msg {
	return messages.FilterMsg{FilterName: m.Title, Action: action}
}
