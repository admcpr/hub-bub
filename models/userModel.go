package models

import (
	"fmt"

	"github.com/admcpr/hub-bub/messages"
	"github.com/admcpr/hub-bub/structs"
	"github.com/cli/go-gh"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UserModel struct {
	Authenticated  bool
	User           structs.User
	SelectedOrgUrl string
	list           list.Model
	loaded         bool
	width          int
	height         int
}

func (m *UserModel) initList() {
	m.list = list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		m.width,
		m.height,
	)
}

func (m UserModel) Init() tea.Cmd {
	return checkLoginStatus
}

func (m UserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:

		var columnStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.HiddenBorder())
		var focusedStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))
		const divisor = 4

		m.height = msg.Height
		m.width = msg.Width

		if !m.loaded {
			columnStyle.Width(msg.Width / divisor)
			focusedStyle.Width(msg.Width / divisor)
			columnStyle.Height(msg.Height - divisor)
			focusedStyle.Height(msg.Height - divisor)
			m.initList()
			m.loaded = true
		}

	case messages.AuthenticationMsg:
		m.Authenticated = true
		m.User = msg.User
		return m, getOrganisations

	case messages.AuthenticationErrorMsg:
		m.Authenticated = false
		return m, nil

	case messages.OrgListMsg:
		m.list = buildOrgListModel(msg.Organisations, m.width, m.height)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			MainModel[UserModelName] = m
			item := m.list.SelectedItem()
			orgModel := &OrganisationModel{
				Title: item.(structs.ListItem).Title(),
				Url:   item.(structs.ListItem).Description(),
			}

			MainModel[OrganisationModelName] = orgModel

			return orgModel, orgModel.GetRepositories
		}
	}

	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m UserModel) View() string {
	if !m.Authenticated {
		return fmt.Sprintln("You are not authenticated try running `gh auth login`. Press q to quit.")
	}

	var docStyle = lipgloss.NewStyle().Margin(1, 2)

	return docStyle.Render(m.list.View())
}

func buildOrgListModel(organisations []structs.Organisation, width, height int) list.Model {
	items := make([]list.Item, len(organisations))
	for i, org := range organisations {
		items[i] = structs.NewListItem(org.Login, org.Url)
	}

	list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Organisations"
	list.SetHeight(height)
	list.SetWidth(width)

	return list
}

func checkLoginStatus() tea.Msg {
	client, err := gh.RESTClient(nil)
	if err != nil {
		return messages.AuthenticationErrorMsg{Err: err}
	}
	response := structs.User{}

	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return messages.AuthenticationErrorMsg{Err: err}
	}

	return messages.AuthenticationMsg{User: response}
}

func getOrganisations() tea.Msg {
	client, err := gh.RESTClient(nil)
	if err != nil {
		return messages.AuthenticationErrorMsg{Err: err}
	}
	response := []structs.Organisation{}

	err = client.Get("user/orgs", &response)
	if err != nil {
		fmt.Println(err)
		return messages.ErrMsg{Err: err}
	}

	return messages.OrgListMsg{Organisations: response}
}
