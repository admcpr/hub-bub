package models

import (
	"hub-bub/structs"

	"github.com/charmbracelet/lipgloss"
)

func RenderTabs(tabSettings []structs.RepositorySettingsTab, width, activeTab int) string {
	inactiveTabBorder := tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder := tabBorderWithBottom("┘", " ", "└")
	inactiveTabStyle := lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(blueLighter).Padding(0)
	activeTabStyle := inactiveTabStyle.Copy().Border(activeTabBorder, true)

	tabs := []string{}
	for _, t := range tabSettings {
		tabs = append(tabs, t.Name)
	}

	tabWidth := ((width) / len(tabs)) - 2

	var renderedTabs []string

	for i, t := range tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(tabs)-1, i == activeTab
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}

		style = style.Border(border).Align(lipgloss.Center)

		if i == activeTab {
			style = style.Foreground(pink)
		}

		if isLast {
			style = style.Width(tabWidth + (width % len(tabs)))
		} else {
			style = style.Width(tabWidth)
		}

		renderedTabs = append(renderedTabs, style.Render(t))
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
