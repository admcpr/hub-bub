package models

import (
	"fmt"

	"github.com/admcpr/hub-bub/messages"
	"github.com/admcpr/hub-bub/structs"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RepositoryModel struct {
	Tabs       []string
	TabContent []string
	TabLists   []structs.RepositorySettingsTab
	ActiveTab  int
	ActiveList list.Model
	width      int
}

func NewRepositoryModel(ornq structs.OrganisationRepositoryNodeQuery, width int) RepositoryModel {
	return RepositoryModel{
		Tabs:       []string{"Overview", "Features", "PRs & Default Branch", "Security", "Wiki", "Settings"},
		TabContent: []string{buildOverview(ornq), "Issues Tab", "Pull Requests Tab", "Projects Tab", "Wiki Tab", "Settings Tab"},
		TabLists:   structs.BuildRepositorySettings(ornq),
		ActiveTab:  0,
		ActiveList: list.New(
			[]list.Item{},
			list.NewDefaultDelegate(),
			width,
			100,
		),
		width: width,
	}
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Center).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

// View implements tea.Model
func (m RepositoryModel) View() string {
	if len(m.Tabs) == 0 {
		return ""
	}

	// doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.Tabs {
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.Tabs)-1, i == m.ActiveTab
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
		style = style.Border(border)
		renderedTabs = append(renderedTabs, style.Render(t))
	}

	// renderer, _ := glamour.NewTermRenderer(
	// 	glamour.WithAutoStyle(),
	// 	glamour.WithWordWrap(m.width),
	// )

	// contentStr, _ := renderer.Render(m.TabContent[m.ActiveTab])

	//row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	return fmt.Sprintf("ActiveTab:%d", m.ActiveTab)

	// m.ActiveList.View()

	// lipgloss.JoinVertical(lipgloss.Top, docStyle.Render(row), docStyle.Render(m.ActiveList.View()))
	// doc.WriteString(row)
	// doc.WriteString("\n")
	// doc.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(m.TabContent[m.ActiveTab]))
	// doc.WriteString(docStyle.Render(m.ActiveList.View()))
	// return docStyle.Render(doc.String())
}

func (m RepositoryModel) Init() tea.Cmd {
	return textarea.Blink
}

func buildTabListModel(tabSettings structs.RepositorySettingsTab, width, height int) list.Model {
	items := make([]list.Item, len(tabSettings.Settings))
	for i, setting := range tabSettings.Settings {
		items[i] = structs.NewListItem(setting.Name, setting.Value)
	}

	list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	// list.Title = fmt.Sprintf("Settings h:%d w:%d", height, width)
	list.SetHeight(height)
	list.SetWidth(width)

	return list
}

func (m RepositoryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case messages.RepositorySelectedMsg:
		m.ActiveList = buildTabListModel(m.TabLists[0], m.width, 100)
		// case tea.KeyMsg:
		// 	switch keypress := msg.String(); keypress {
		// 	case "ctrl+c", "q":
		// 		return m, tea.Quit
		// 	case "right", "l", "n", "tab":
		// 		m.activeTab = min(m.activeTab+1, len(m.Tabs)-1)
		// 		return m, nil
		// 	case "left", "h", "p", "shift+tab":
		// 		m.activeTab = max(m.activeTab-1, 0)
		// 		return m, nil
		// 	}
	}

	return m, nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
