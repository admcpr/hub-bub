package main

import (
	"fmt"
	"os"

	"github.com/admcpr/hub-bub/models"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

func checkLoginStatus() tea.Msg {
	// Use an API helper to grab repository tags
	client, err := gh.RESTClient(nil)
	if err != nil {
		return models.AuthenticationErrorMsg{Err: err}
	}
	response := models.User{}

	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return models.AuthenticationErrorMsg{Err: err}
	}

	return models.AuthenticationMsg{User: response}
}

func getOrganisations() tea.Msg {
	client, err := gh.RESTClient(nil)
	if err != nil {
		return models.AuthenticationErrorMsg{Err: err}
	}
	response := []models.Organisation{}

	err = client.Get("user/orgs", &response)
	if err != nil {
		fmt.Println(err)
		return models.ErrMsg{Err: err}
	}

	return models.OrgListMsg{Organisations: response}
}

func getRepositories() tea.Msg {
	client, err := gh.RESTClient(nil)
	if err != nil {
		return models.AuthenticationErrorMsg{Err: err}
	}
	response := []models.Repository{}

	// err = client.Get(fmt.Sprintf("user/%s/repos", m.OrganisationTable.SelectedRow()[0]), &response)
	err = client.Get(fmt.Sprintf("user/%s/repos", "?"), &response)
	if err != nil {
		fmt.Println(err)
		return models.ErrMsg{Err: err}
	}

	return models.RepositoryListMsg{Repositories: response}
}

func initialModel() model {
	return model{}
}

func buildOrganisationTable(organisations []models.Organisation) table.Model {
	columns := []table.Column{
		{Title: "Login", Width: 20},
		{Title: "Url", Width: 80},
	}

	rows := make([]table.Row, len(organisations))
	for i, org := range organisations {
		rows[i] = table.Row{org.Login, org.Url}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(organisations)),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return t
}

type model struct {
	Authenticated     bool
	User              models.User
	SelectedOrgUrl    string
	OrganisationTable table.Model
	RepositoryTable   table.Model
}

func (m model) Init() tea.Cmd {
	return checkLoginStatus
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case models.AuthenticationMsg:
		m.Authenticated = true
		m.User = msg.User
		return m, getOrganisations

	case models.AuthenticationErrorMsg:
		m.Authenticated = false
		return m, nil

	case models.OrgListMsg:
		m.OrganisationTable = buildOrganisationTable(msg.Organisations)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter", " ":
			// return m, getRepositories
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.OrganisationTable.SelectedRow()[1]),
			)
		}
	}

	m.OrganisationTable, cmd = m.OrganisationTable.Update(msg)

	return m, cmd
}

func (m model) View() string {
	s := fmt.Sprintln("Press q to quit.")

	if m.Authenticated {
		s += fmt.Sprintf("Hello %s\n", m.User.Name)
	} else {
		return fmt.Sprintln("You are not authenticated try running `gh auth login`")
	}

	// if (m.OrganisationTable != table.Model{}) {
	// 	s += baseStyle.Render(m.OrganisationTable.View()) + "\n"
	// }

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
