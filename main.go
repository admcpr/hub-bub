package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
)

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

func YesNo(b bool) string {
	if b {
		return "Yes"
	}
	return "No"
}

func buildRepositoryTable(repositories []Repository) table.Model {
	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Issues", Width: 10},
		{Title: "Wiki", Width: 10},
		{Title: "Projects", Width: 10},
		{Title: "Rebase Merge", Width: 10},
		{Title: "Auto Merge", Width: 10},
		{Title: "Delete Branch On Merge", Width: 10},
	}

	rows := make([]table.Row, len(repositories))
	for i, repo := range repositories {
		rows[i] = table.Row{
			repo.Name,
			YesNo(repo.HasIssues),
			YesNo(repo.HasWiki),
			YesNo(repo.HasProjects),
			YesNo(repo.AllowRebaseMerge),
			YesNo(repo.AllowAutoMerge),
			YesNo(repo.DeleteBranchOnMerge),
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(repositories)),
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
	models = []tea.Model{&UserModel{}, &OrganisationModel{}}

	p := tea.NewProgram(models[user])
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
