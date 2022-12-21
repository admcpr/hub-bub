package models

import (
	"log"

	"github.com/admcpr/hub-bub/messages"
	"github.com/admcpr/hub-bub/queries"
	"github.com/admcpr/hub-bub/structs"
	"github.com/admcpr/hub-bub/utils"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/cli/go-gh"
	graphql "github.com/cli/shurcooL-graphql"
)

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

	case messages.RepositoryListMsg:
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
			return MainModel[UserModelName], nil
		}
	}

	m.RepositoryTable, cmd = m.RepositoryTable.Update(msg)

	return m, cmd
}

// View implements tea.Model
func (m OrganisationModel) View() string {
	return utils.BaseStyle.Render(m.RepositoryTable.View()) + "\n"
}

func (m OrganisationModel) GetRepositories() tea.Msg {
	client, err := gh.GQLClient(nil)
	if err != nil {
		return messages.AuthenticationErrorMsg{Err: err}
	}

	var organizationQuery = queries.OrganizationQuery{}

	variables := map[string]interface{}{
		"login": graphql.String(m.Title),
		"first": graphql.Int(30),
	}
	err = client.Query("OrganizationRepositories", &organizationQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	return messages.RepositoryListMsg{OrganizationQuery: organizationQuery}
}

func buildOrganisationTable(organisations []structs.Organisation) table.Model {
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

func buildRepositoryTable(organizationQuery queries.OrganizationQuery) table.Model {
	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Issues", Width: 10},
		{Title: "Wiki", Width: 10},
		{Title: "Projects", Width: 10},
		{Title: "Rebase Merge", Width: 10},
		{Title: "Auto Merge", Width: 10},
		{Title: "Delete Branch On Merge", Width: 10},
	}

	edges := organizationQuery.Organization.Repositories.Edges

	rows := make([]table.Row, len(edges))
	for i, repo := range edges {
		rows[i] = table.Row{
			repo.Node.Name,
			utils.YesNo(repo.Node.HasIssuesEnabled),
			utils.YesNo(repo.Node.HasWikiEnabled),
			utils.YesNo(repo.Node.HasProjectsEnabled),
			utils.YesNo(repo.Node.RebaseMergeAllowed),
			utils.YesNo(repo.Node.AutoMergeAllowed),
			utils.YesNo(repo.Node.DeleteBranchOnMerge),
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(len(edges)),
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
