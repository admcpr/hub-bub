package models

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	// purple     = "#cdb4db"
	// pink       = "#ffc8dd"
	// pinkDarker = "#ffafcc"
	// blue       = "#bde0fe"
	blueDarker = "#a2d2ff"
	white      = "#FFFDF5"

	appStyle = lipgloss.NewStyle().Padding(1, 0)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(white)).
			Background(lipgloss.Color(blueDarker)).
			Padding(0, 1)

	titleHeight       = 2
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	borderColor       = lipgloss.Color(blueDarker)
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(borderColor).Padding(0)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)

	settingsStyle = appStyle.Copy().Border(settingsBorder()).
			BorderForeground(borderColor).Padding(0).Margin(0)
)

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

func settingsBorder() lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.Top = ""
	border.TopLeft = "│"
	border.TopRight = "│"
	return border
}
