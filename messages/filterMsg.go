package messages

import "hub-bub/structs"

type FilterMsg struct {
	Filter structs.SettingGetter
	Action structs.RepositoryFilterAction
}