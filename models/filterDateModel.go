package models

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterDateModel struct {
	Title     string
	fromInput textinput.Model
	toInput   textinput.Model
}

func (m FilterDateModel) Init() tea.Cmd {
	return nil
}

func (m FilterDateModel) Update(msg tea.Msg) (FilterDateModel, tea.Cmd) {
	return m, nil
}

func (m FilterDateModel) View() string {
	return m.Title + "\n\n" + m.fromInput.View() + "\n\n" + m.toInput.View()

}
