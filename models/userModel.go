package models

import (
	"fmt"

	"github.com/admcpr/hub-bub/messages"
	"github.com/admcpr/hub-bub/structs"
	"github.com/cli/go-gh"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type UserModel struct {
	Authenticated     bool
	User              structs.User
	SelectedOrgUrl    string
	OrganisationTable table.Model
	list              list.Model
	loaded            bool
}

func (m *UserModel) initList(width, height int) {
	m.list = list.New(
		[]list.Item{},
		list.NewDefaultDelegate(),
		100,
		100,
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

		if !m.loaded {
			columnStyle.Width(msg.Width / divisor)
			focusedStyle.Width(msg.Width / divisor)
			columnStyle.Height(msg.Height - divisor)
			focusedStyle.Height(msg.Height - divisor)
			m.initList(msg.Width, msg.Height)
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
		// m.OrganisationTable = buildOrganisationTable(msg.Organisations)
		m.list = buildListModel(msg.Organisations)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			MainModel[UserModelName] = m
			orgModel := &OrganisationModel{
				Title: m.OrganisationTable.SelectedRow()[0],
				Url:   m.OrganisationTable.SelectedRow()[1],
			}

			MainModel[OrganisationModelName] = orgModel

			return orgModel, orgModel.GetRepositories
		}
	}

	// m.OrganisationTable, cmd = m.OrganisationTable.Update(msg)
	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m UserModel) View() string {
	s := fmt.Sprintln("Press q to quit.")

	if !m.Authenticated {
		return fmt.Sprintln("You are not authenticated try running `gh auth login`")
	}

	s += fmt.Sprintf("Hello %s, press Enter to select an organisation.\n", m.User.Name)
	// s += utils.BaseStyle.Render(m.OrganisationTable.View()) + "\n"

	var docStyle = lipgloss.NewStyle().Margin(1, 2)

	s += docStyle.Render(m.list.View())

	return s
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func buildListModel(organisations []structs.Organisation) list.Model {
	items := make([]list.Item, len(organisations))
	for i, org := range organisations {
		items[i] = item{title: org.Login, desc: org.Url}
	}

	list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Organisations"
	list.SetHeight(3)
	list.SetWidth(100)

	return list
}

func checkLoginStatus() tea.Msg {
	// Use an API helper to grab repository tags
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
