package models

import (
	"fmt"
	"io"

	"github.com/admcpr/hub-bub/structs"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO: Pull this out into it's own file
type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 1 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(structs.ListItem)
	if !ok {
		return
	}

	// str := fmt.Sprintf("%s > %s", i.Title(), i.Description())

	statusNugget := lipgloss.NewStyle().
		Foreground(lipgloss.Color(white)).
		Padding(0, 1)

	// statusBarStyle := lipgloss.NewStyle().
	// 	Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
	// 	Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	// statusStyle := lipgloss.NewStyle().
	// 	Inherit(statusBarStyle).
	// 	Foreground(lipgloss.Color("#FFFDF5")).
	// 	Background(lipgloss.Color("#FF5F87")).
	// 	Padding(0, 1).
	// 	MarginRight(1)

	encodingStyle := statusNugget.Copy().
		Background(lipgloss.Color(pink)).
		Align(lipgloss.Right)

	// str := lipgloss.JoinHorizontal(lipgloss.Left, i.Title(), i.Description())

	title := statusNugget.Render(i.Title())
	description := encodingStyle.Render(i.Description())

	// fn := itemStyle.Render
	// if index == m.Index() {
	// 	fn = func(s string) string {
	// 		return selectedItemStyle.Render("> " + s)
	// 	}
	// }
	fn := lipgloss.JoinHorizontal

	fmt.Fprint(w, fn(lipgloss.Left, title, description))
}
