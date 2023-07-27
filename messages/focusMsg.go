package messages

import "hub-bub/consts"

type FocusMsg struct {
	Focus consts.Focus
}

func NewFocusMsg(focus consts.Focus) FocusMsg {
	return FocusMsg{Focus: focus}
}
