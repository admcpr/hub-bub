package models

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterNumberModel struct {
	Title     string
	fromInput textinput.Model
	toInput   textinput.Model
}

func numberValidator(s, prompt string) error {
	_, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("please enter a number for the `%s` value", prompt)
	}

	return nil
}

func NewFilterNumberModel(title string, from, to int) FilterNumberModel {
	m := FilterNumberModel{
		Title:     title,
		fromInput: textinput.New(),
		toInput:   textinput.New(),
	}

	m.fromInput.Placeholder = fmt.Sprint(from)
	m.fromInput.Prompt = "From: "
	m.fromInput.CharLimit = 4
	m.fromInput.Validate = func(s string) error { return numberValidator(s, m.fromInput.Prompt) }

	m.toInput.Placeholder = fmt.Sprint(to)
	m.toInput.Prompt = "To: "
	m.toInput.CharLimit = 4
	m.toInput.Validate = func(s string) error { return numberValidator(s, m.toInput.Prompt) }

	m.fromInput.Focus()

	return m
}

func (m FilterNumberModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m FilterNumberModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
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

func (m FilterNumberModel) View() string {
	return m.Title + " " + m.fromInput.View() + " " + m.toInput.View()
}

func (m *FilterNumberModel) GetValue() (int, int) {
	from, _ := strconv.Atoi(m.fromInput.Value())
	to, _ := strconv.Atoi(m.toInput.Value())
	return from, to
}
