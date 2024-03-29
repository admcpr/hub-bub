package style

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

var (
	Pink         = lipgloss.Color("#f72585")
	Purple       = lipgloss.Color("#7209b7")
	PurpleDarker = lipgloss.Color("#3a0ca3")
	Blue         = lipgloss.Color("#4361ee")
	BlueLighter  = lipgloss.Color("#4cc9f0")
	White        = lipgloss.Color("#dddddd")

	AppStyle = lipgloss.NewStyle().Padding(0, 0).Foreground(White).BorderForeground(Blue)

	TitleStyle = lipgloss.NewStyle().
			Foreground(Blue).
			BorderForeground(BlueLighter).
			Border(lipgloss.RoundedBorder(), true).
			Padding(0, 1)

	ErrorStyle = lipgloss.NewStyle().Foreground(PurpleDarker)

	DefaultDelegate = BuildDefaultDelegate()
)

func BuildDefaultDelegate() list.DefaultDelegate {
	defaultDelegate := list.NewDefaultDelegate()
	defaultDelegate.Styles.SelectedTitle.Foreground(Pink)
	defaultDelegate.Styles.SelectedTitle.BorderForeground(Pink)
	defaultDelegate.Styles.SelectedDesc.Foreground(Purple)
	defaultDelegate.Styles.SelectedDesc.BorderForeground(Pink)

	return defaultDelegate
}
