package models

type Focus int

const (
	focusList Focus = iota
	focusTabs
	focusFilter
)

func (f Focus) Next() Focus {
	switch f {
	case focusList:
		return focusTabs
	default:
		return focusFilter
	}
}

func (f Focus) Prev() Focus {
	switch f {
	case focusFilter:
		return focusTabs
	default:
		return focusList
	}
}
