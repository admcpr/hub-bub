package main

import (
	"fmt"
	"os"

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
		return AuthenticationErrorMsg{Err: err}
	}
	response := User{}

	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return AuthenticationErrorMsg{Err: err}
	}

	return AuthenticationMsg{User: response}
}

func getOrganisations() tea.Msg {
	client, err := gh.RESTClient(nil)
	if err != nil {
		return AuthenticationErrorMsg{Err: err}
	}
	response := []Organisation{}

	err = client.Get("user/orgs", &response)
	if err != nil {
		fmt.Println(err)
		return ErrMsg{Err: err}
	}

	return OrgListMsg{Organisations: response}
}

func getRepositories() tea.Msg {
	client, err := gh.RESTClient(nil)
	if err != nil {
		return AuthenticationErrorMsg{Err: err}
	}
	response := []Repository{}

	// err = client.Get(fmt.Sprintf("user/%s/repos", m.OrganisationTable.SelectedRow()[0]), &response)
	err = client.Get(fmt.Sprintf("user/%s/repos", "?"), &response)
	if err != nil {
		fmt.Println(err)
		return ErrMsg{Err: err}
	}

	return RepositoryListMsg{Repositories: response}
}

func initialModel() UserModel {
	return UserModel{}
}

func buildOrganisationTable(organisations []Organisation) table.Model {
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

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
