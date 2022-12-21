package models

import tea "github.com/charmbracelet/bubbletea"

/*
Model management
Need to replace this with the a MainModel, see nested models youtube
*/
type modelName int

var MainModel []tea.Model

const (
	UserModelName modelName = iota
	OrganisationModelName
)
