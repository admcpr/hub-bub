package models

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterDateModel struct {
	Title     string
	fromInput textinput.Model
	toInput   textinput.Model
}

func NewFilterDateModel(title string, from, to time.Time) FilterDateModel {
	m := FilterDateModel{
		Title:     title,
		fromInput: textinput.New(),
		toInput:   textinput.New(),
	}

	m.fromInput.SetValue(from.Format("2006-01-02"))
	m.toInput.SetValue(to.Format("2006-01-02"))

	return m
}

func (m FilterDateModel) Init() tea.Cmd {
	return m.Focus()
}

func (m FilterDateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m FilterDateModel) View() string {
	return m.Title + " between " + m.fromInput.View() + " and " + m.toInput.View()
}

func (m *FilterDateModel) Focus() tea.Cmd {
	return m.fromInput.Focus()
}

func (m *FilterDateModel) GetValue() (time.Time, time.Time) {
	from, _ := time.Parse("2006-01-02", m.fromInput.Value())
	to, _ := time.Parse("2006-01-02", m.toInput.Value())
	return from, to
}
