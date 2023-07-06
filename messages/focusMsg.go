package messages

import "hub-bub/consts"

type FocusMessage struct {
	Focus consts.Focus
}

func NewFocusMsg(focus consts.Focus) FocusMessage {
	return FocusMessage{Focus: focus}
}
