package models

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterNumberModel struct {
	Title     string
	from      int
	to        int
	fromInput textinput.Model
	toInput   textinput.Model
}

func NewFilterNumberModel(from, to int) FilterNumberModel {
	return FilterNumberModel{from: from, to: to}
}

func (m FilterNumberModel) Init() tea.Cmd {
	return nil
}

func (m FilterNumberModel) Update(msg tea.Msg) (FilterNumberModel, tea.Cmd) {
	return m, nil
}

func (m FilterNumberModel) View() string {
	return m.Title + "\n\n" + m.fromInput.View() + "\n\n" + m.toInput.View()
}
