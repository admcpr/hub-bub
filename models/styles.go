package models

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Colors
	purple     = lipgloss.Color("#cdb4db")
	pink       = lipgloss.Color("#ffc8dd")
	pinkDarker = lipgloss.Color("#ffafcc")
	blue       = lipgloss.Color("#bde0fe")
	blueDarker = lipgloss.Color("#a2d2ff")
	// white      = lipgloss.Color("#FFFDF5")

	appStyle = lipgloss.NewStyle().Padding(0, 0).Foreground(purple).BorderForeground(blue)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(blueDarker)).
			Border(lipgloss.RoundedBorder(), true).
			Padding(0, 1)

	defaultDelegate = buildDefaultDelegate()
)

func buildDefaultDelegate() list.DefaultDelegate {
	defaultDelegate := list.NewDefaultDelegate()
	defaultDelegate.Styles.SelectedTitle.Foreground(blueDarker)
	defaultDelegate.Styles.SelectedTitle.BorderForeground(blueDarker)
	defaultDelegate.Styles.SelectedDesc.Foreground(blueDarker)
	defaultDelegate.Styles.SelectedDesc.BorderForeground(blueDarker)

	return defaultDelegate
}
