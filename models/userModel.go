package models

import (
	"fmt"

	"github.com/admcpr/hub-bub/messages"
	"github.com/admcpr/hub-bub/structs"
	"github.com/admcpr/hub-bub/utils"
	"github.com/cli/go-gh"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type UserModel struct {
	Authenticated     bool
	User              structs.User
	SelectedOrgUrl    string
	OrganisationTable table.Model
}

func (m UserModel) Init() tea.Cmd {
	return checkLoginStatus
}

func (m UserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case messages.AuthenticationMsg:
		m.Authenticated = true
		m.User = msg.User
		return m, getOrganisations

	case messages.AuthenticationErrorMsg:
		m.Authenticated = false
		return m, nil

	case messages.OrgListMsg:
		m.OrganisationTable = buildOrganisationTable(msg.Organisations)
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

	m.OrganisationTable, cmd = m.OrganisationTable.Update(msg)

	return m, cmd
}

func (m UserModel) View() string {
	s := fmt.Sprintln("Press q to quit.")

	if !m.Authenticated {
		return fmt.Sprintln("You are not authenticated try running `gh auth login`")
	}

	s += fmt.Sprintf("Hello %s, press Enter to select an organisation.\n", m.User.Name)
	s += utils.BaseStyle.Render(m.OrganisationTable.View()) + "\n"

	return s
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
