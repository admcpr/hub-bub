package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/cli/go-gh"
)

type user struct {
	Name  string
	Login string
}

type organisation struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	Repositories []repository
}

type repository struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func checkLoginStatus() tea.Msg {
	// Use an API helper to grab repository tags
	client, err := gh.RESTClient(nil)
	if err != nil {
		return authenticationErrorMsg{err: err}
	}
	response := user{}

	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return authenticationErrorMsg{err: err}
	}

	return authenticationMsg{User: response}
}

type authenticationErrorMsg struct {
	err error
}

type authenticationMsg struct {
	User user
}

type errMsg struct{ err error }

type model struct {
	Authenticated bool
	User          user
	cursor        int
	organisations []organisation
	selected      map[int]struct{}
}

func initialModel() model {
	return model{
		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return checkLoginStatus
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case authenticationMsg:
		m.Authenticated = true
		m.User = msg.User
		return m, nil

	case authenticationErrorMsg:
		m.Authenticated = false
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.organisations)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintln("Press q to quit.")

	if m.Authenticated {
		s += fmt.Sprintf("Hello %s", m.User.Name)
	} else {
		s += fmt.Sprintln("You are not authenticated try running `gh auth login`")
	}

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
