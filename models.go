package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
)

/* Model management */
type modelName int

var models []tea.Model

const (
	user modelName = iota
	organisation
)

/* User model */
type UserModel struct {
	Authenticated     bool
	User              User
	SelectedOrgUrl    string
	OrganisationTable table.Model
}

func (m UserModel) Init() tea.Cmd {
	return checkLoginStatus
}

func (m UserModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case AuthenticationMsg:
		m.Authenticated = true
		m.User = msg.User
		return m, getOrganisations

	case AuthenticationErrorMsg:
		m.Authenticated = false
		return m, nil

	case OrgListMsg:
		m.OrganisationTable = buildOrganisationTable(msg.Organisations)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			models[user] = m
			orgModel := &OrganisationModel{
				Title: m.OrganisationTable.SelectedRow()[0],
				Url:   m.OrganisationTable.SelectedRow()[1],
			}

			models[organisation] = orgModel

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
	s += baseStyle.Render(m.OrganisationTable.View()) + "\n"

	return s
}

/* Repository model */
type OrganisationModel struct {
	Title           string
	Url             string
	RepositoryTable table.Model
}

func (m OrganisationModel) Init() tea.Cmd {
	return nil
}

func (m OrganisationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case RepositoryListMsg:
		// m.RepositoryTable = buildRepositoryTable(msg.Repositories)
		m.RepositoryTable = buildRepositoryTable(msg.OrganizationQuery)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.RepositoryTable.SelectedRow()[1]),
			)
		case "esc":
			return models[user], nil
		}
	}

	m.RepositoryTable, cmd = m.RepositoryTable.Update(msg)

	return m, cmd
}

// View implements tea.Model
func (m OrganisationModel) View() string {
	return baseStyle.Render(m.RepositoryTable.View()) + "\n"
}

func (m OrganisationModel) GetRepositories() tea.Msg {
	client, err := gh.GQLClient(nil)
	if err != nil {
		return AuthenticationErrorMsg{Err: err}
	}

	var organizationQuery = OrganizationQuery{}

	variables := map[string]interface{}{
		"login": graphql.String(m.Title),
		"first": graphql.Int(30),
	}
	err = client.Query("OrganizationRepositories", &organizationQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	return RepositoryListMsg{OrganizationQuery: organizationQuery}
}
