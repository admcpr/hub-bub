package consts

type FilterAction int

const (
	FilterConfirm FilterAction = iota
	FilterDelete
	FilterCancel
)
