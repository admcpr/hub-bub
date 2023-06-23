package models

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	pink         = lipgloss.Color("#f72585")
	purple       = lipgloss.Color("#7209b7")
	purpleDarker = lipgloss.Color("#3a0ca3")
	blue         = lipgloss.Color("#4361ee")
	blueLighter  = lipgloss.Color("#4cc9f0")
	white        = lipgloss.Color("#dddddd")

	appStyle = lipgloss.NewStyle().Padding(0, 0).Foreground(white).BorderForeground(blue)

	titleStyle = lipgloss.NewStyle().
			Foreground(blue).
			BorderForeground(blueLighter).
			Border(lipgloss.RoundedBorder(), true).
			Padding(0, 1)

	errorStyle = lipgloss.NewStyle().Foreground(purpleDarker)

	defaultDelegate = buildDefaultDelegate()
)

func buildDefaultDelegate() list.DefaultDelegate {
	defaultDelegate := list.NewDefaultDelegate()
	defaultDelegate.Styles.SelectedTitle.Foreground(pink)
	defaultDelegate.Styles.SelectedTitle.BorderForeground(pink)
	defaultDelegate.Styles.SelectedDesc.Foreground(purple)
	defaultDelegate.Styles.SelectedDesc.BorderForeground(pink)

	return defaultDelegate
}
