package models

import (
	"fmt"
	"hub-bub/messages"
	"hub-bub/structs"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type FilterModel struct {
	RepositoryFilter structs.RepositoryFilter
	textinput        textinput.Model
}

func NewFilterModel(tab string, setting structs.SettingGetter) FilterModel {
	filterModel := FilterModel{
		RepositoryFilter: structs.RepositoryFilter{Tab: tab, SettingGetter: setting},
		textinput:        textinput.New(),
	}
	filterModel.textinput.Placeholder = setting.GetValue()
	return filterModel
}

func (m FilterModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m FilterModel) Update(msg tea.Msg) (FilterModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m, m.SendMsg(structs.CancelAction)
		}
	}

	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

func (m FilterModel) View() string {
	rf := m.RepositoryFilter
	// switch rf.SettingGetter.GetType().String() {
	// case "bool":
	// 	return fmt.Sprintf("bool setting '%s' on tab '%s' with value '%s'", rf.SettingGetter.GetName(), rf.Tab, rf.SettingGetter.GetValue())
	// case "int":
	// 	return fmt.Sprintf("int setting '%s' on tab '%s' with value '%s'", rf.SettingGetter.GetName(), rf.Tab, rf.SettingGetter.GetValue())
	// default:
	// 	return fmt.Sprintf("date setting '%s' on tab '%s' with value '%s'", rf.SettingGetter.GetName(), rf.Tab, rf.SettingGetter.GetValue())
	// }

	return fmt.Sprintf("%s\n\n%s", rf.SettingGetter.GetName(), m.textinput.View())
}

func (m FilterModel) SendMsg(action structs.RepositoryFilterAction) tea.Cmd {
	return func() tea.Msg {
		return messages.FilterMsg{Filter: m.RepositoryFilter.SettingGetter, Action: action}
	}
}
