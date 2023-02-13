package models

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	purple     = "#cdb4db"
	pink       = "#ffc8dd"
	pinkDarker = "#ffafcc"
	blue       = "#bde0fe"
	blueDarker = "#a2d2ff"
	white      = "#FFFDF5"

	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(white)).
			Background(lipgloss.Color(blueDarker)).
			Padding(0, 1)

	titleHeight = 2

	// Custom list
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(blue))

	// Tabs
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.Color(pinkDarker)
	borderColor       = lipgloss.Color(blueDarker)
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(borderColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()

	settingsStyle = appStyle.Copy().Border(lipgloss.RoundedBorder(), false, true, true).BorderForeground(borderColor).Padding(0, 1, 1, 1)
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}
