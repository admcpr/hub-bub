package models

import (
	"fmt"
	"hub-bub/messages"
	"hub-bub/structs"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterIntModel struct {
	Tab       string
	Title     string
	fromInput textinput.Model
	toInput   textinput.Model
}

func intValidator(s, prompt string) error {
	_, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("please enter an integer for the `%s` value", prompt)
	}

	return nil
}

func NewFilterIntModel(tab, title string, from, to int) FilterIntModel {
	m := FilterIntModel{
		Tab:       tab,
		Title:     title,
		fromInput: textinput.New(),
		toInput:   textinput.New(),
	}

	m.fromInput.Placeholder = fmt.Sprint(from)
	m.fromInput.Prompt = "From: "
	m.fromInput.CharLimit = 4
	m.fromInput.Validate = func(s string) error { return intValidator(s, m.fromInput.Prompt) }

	m.toInput.Placeholder = fmt.Sprint(to)
	m.toInput.Prompt = "To: "
	m.toInput.CharLimit = 4
	m.toInput.Validate = func(s string) error { return intValidator(s, m.toInput.Prompt) }

	m.fromInput.Focus()

	return m
}

func (m FilterIntModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m FilterIntModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyEsc.String():
			return m, m.Cancel
		case tea.KeyTab.String():
			if m.fromInput.Focused() {
				m.fromInput.Blur()
				m.toInput.Focus()
			} else {
				m.toInput.Blur()
				m.fromInput.Focus()
			}
		default:
			if m.fromInput.Focused() {
				m.fromInput, cmd = m.fromInput.Update(msg)
			} else {
				m.toInput, cmd = m.toInput.Update(msg)
			}
		}
	}

	return m, cmd
}

func (m FilterIntModel) View() string {
	return m.Title + " " + m.fromInput.View() + " " + m.toInput.View()
}

func (m *FilterIntModel) GetValue() (int, int) {
	from, _ := strconv.Atoi(m.fromInput.Value())
	to, _ := strconv.Atoi(m.toInput.Value())
	return from, to
}

func (m FilterIntModel) Cancel() tea.Msg {
	return messages.NewCancelFilterMsg(structs.NewFilterInt(m.Tab, m.Title, 0, 0))
}

func (m FilterIntModel) Confirm() tea.Msg {
	from, to := m.GetValue()

	return messages.NewAddFilterMsg(structs.NewFilterInt(m.Tab, m.Title, from, to))
}
