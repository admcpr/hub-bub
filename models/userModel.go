package models

import (
	"fmt"

	"hub-bub/messages"
	"hub-bub/structs"

	"github.com/cli/go-gh"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
		m.height = msg.Height
		m.width = msg.Width

		if !m.loaded {
			m.initList()
			m.loaded = true
		}
		return m, nil

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
			orgModel := NewOrgModel(item.(structs.ListItem).Title(), m.width, m.height)

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

	return appStyle.Render(m.list.View())
}

func buildOrgListModel(organisations []structs.Organisation, width, height int) list.Model {
	items := make([]list.Item, len(organisations))
	for i, org := range organisations {
		items[i] = structs.NewListItem(org.Login, org.Url)
	}

	list := list.New(items, list.NewDefaultDelegate(), width, height-titleHeight)

	list.Title = "Organisations"
	list.SetStatusBarItemName("Organisation", "Organisations")
	list.Styles.Title = titleStyle
	list.SetShowTitle(true)

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
