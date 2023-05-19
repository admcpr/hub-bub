package main

import (
	"fmt"
	"os"

	"hub-bub/models"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	models.MainModel = []tea.Model{models.NewUserModel(), &models.OrgModel{}}

	p := tea.NewProgram(models.MainModel[models.UserModelName])
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
