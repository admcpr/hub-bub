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

func NewFilterNumberModel(title string, from, to int) FilterNumberModel {
	m := FilterNumberModel{
		Title:     title,
		fromInput: textinput.New(),
		toInput:   textinput.New(),
	}

	m.fromInput.SetValue(fmt.Sprint(from))
	m.toInput.SetValue(fmt.Sprint(to))

	return m
}

func (m FilterNumberModel) Init() tea.Cmd {
	return m.Focus()
}

func (m FilterNumberModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m FilterNumberModel) View() string {
	return m.Title + " between " + m.fromInput.View() + " and " + m.toInput.View()
}

func (m *FilterNumberModel) Focus() tea.Cmd {
	return m.fromInput.Focus()
}

func (m *FilterNumberModel) GetValue() (int, int) {
	from, _ := strconv.Atoi(m.fromInput.Value())
	to, _ := strconv.Atoi(m.toInput.Value())
	return from, to
}
