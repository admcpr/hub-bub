package messages

import "hub-bub/structs"

type FilterAction int

const (
	FilterConfirm FilterAction = iota
	FilterDelete
	FilterCancel
)

type FilterMsg struct {
	Action FilterAction
	Filter structs.Filter
}

func NewConfirmFilterMsg(filter structs.Filter) FilterMsg {
	return FilterMsg{Action: FilterConfirm, Filter: filter}
}

func NewDeleteFilterMsg(filter structs.Filter) FilterMsg {
	return FilterMsg{Action: FilterDelete, Filter: filter}
}

func NewCancelFilterMsg(filter structs.Filter) FilterMsg {
	return FilterMsg{Action: FilterCancel, Filter: filter}
}
