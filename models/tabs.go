package models

import (
	"hub-bub/structs"
	"hub-bub/style"

	"github.com/charmbracelet/lipgloss"
)

func RenderTabs(tabSettings []structs.SettingsTab, width, activeTab int) string {
	inactiveTabBorder := tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder := tabBorderWithBottom("┘", " ", "└")
	inactiveTabStyle := lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(style.BlueLighter).Padding(0)
	activeTabStyle := inactiveTabStyle.Copy().Border(activeTabBorder, true)

	tabs := []string{}
	for _, t := range tabSettings {
		tabs = append(tabs, t.Name)
	}

	tabWidth := ((width) / len(tabs)) - 2

	var renderedTabs []string

	for i, t := range tabs {
		var tabStyle lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(tabs)-1, i == activeTab
		if isActive {
			tabStyle = activeTabStyle.Copy()
		} else {
			tabStyle = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := tabStyle.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}

		tabStyle = tabStyle.Border(border).Align(lipgloss.Center).MaxHeight(3)

		if i == activeTab {
			tabStyle = tabStyle.Foreground(style.Pink)
		}

		if isLast {
			tabStyle = tabStyle.Width(tabWidth + (width % len(tabs)))
		} else {
			tabStyle = tabStyle.Width(tabWidth)
		}

		renderedTabs = append(renderedTabs, tabStyle.Render(t))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	return row
}

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
