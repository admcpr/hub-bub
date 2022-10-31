package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
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
			models[organisation] = OrganisationModel{
				Title: m.OrganisationTable.SelectedRow()[0],
				Url:   m.OrganisationTable.SelectedRow()[1],
			}

			return models[organisation], nil
			// var cmd tea.Cmd
			// cmd = getRepositories(m.OrganisationTable.SelectedRow()[0])

			// return m, getRepositories
			// return m, tea.Batch(
			// 	tea.Printf("Let's go to %s!", m.OrganisationTable.SelectedRow()[1]),4
			// )
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
	return getRepositories
}

func (m OrganisationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case RepositoryListMsg:
		// m.RepositoryTable = buildRepositoryTable(msg.Repositories)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.RepositoryTable.SelectedRow()[1]),
			)
		}
	}

	m.RepositoryTable, cmd = m.RepositoryTable.Update(msg)

	return m, cmd
}

// View implements tea.Model
func (OrganisationModel) View() string {
	panic("unimplemented")
}
