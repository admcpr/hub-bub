package messages

import "hub-bub/structs"

type FilterMsg struct {
	Filter structs.RepositoryFilter
	Action structs.RepositoryFilterAction
}
