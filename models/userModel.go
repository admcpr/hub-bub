package models

import (
	"fmt"

	"hub-bub/messages"
	"hub-bub/structs"

	"github.com/cli/go-gh"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type UserModel struct {
	Authenticating bool
	Authenticated  bool
	User           structs.User
	SelectedOrgUrl string
	list           list.Model
	loaded         bool
	width          int
	height         int
	spinner        spinner.Model
}

func NewUserModel() UserModel {
	return UserModel{
		Authenticating: true,
		list: list.New(
			[]list.Item{},
			list.NewDefaultDelegate(),
			0,
			0,
		),
		spinner: spinner.New(spinner.WithSpinner(spinner.Pulse)),
	}
}

func (m UserModel) Init() tea.Cmd {
	return tea.Batch(checkLoginStatus, m.spinner.Tick)
}

func (m UserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		if !m.loaded {
			m.list.SetWidth(m.width)
			m.list.SetHeight(m.height)
			m.loaded = true
		}
		return m, nil

	case messages.AuthenticationMsg:
		m.Authenticating = false
		m.Authenticated = true
		m.User = msg.User
		return m, getOrganisations

	case messages.AuthenticationErrorMsg:
		m.Authenticating = false
		m.Authenticated = false
		return m, nil

	case messages.OrgListMsg:
		m.list = buildOrgListModel(msg.Organisations, m.width, m.height, m.User)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter", " ":
			MainModel[UserModelName] = m
			item := m.list.SelectedItem()
			orgModel := NewOrgModel(item.(structs.ListItem).Title(), m.width, m.height)

			MainModel[OrganisationModelName] = orgModel

			return orgModel, orgModel.GetRepositories
		}

	default:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	m.list, cmd = m.list.Update(msg)

	return m, cmd
}

func (m UserModel) View() string {
	if !m.Authenticating && !m.Authenticated {
		return fmt.Sprintln("You are not authenticated try running `gh auth login`. Press q to quit.")
	}
	if m.Authenticating {
		return fmt.Sprintf("%s Authenticating with github", m.spinner.View())
	}

	return appStyle.Render(m.list.View())
}

func buildOrgListModel(organisations []structs.Organisation, width, height int, user structs.User) list.Model {
	items := make([]list.Item, len(organisations))
	for i, org := range organisations {
		items[i] = structs.NewListItem(org.Login, org.Url)
	}

	list := list.New(items, defaultDelegate, width, height-2)

	list.Title = "User: " + user.Name
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
