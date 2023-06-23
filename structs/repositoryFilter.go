package structs

type RepositoryFilter struct {
	Tab           string
	SettingGetter SettingGetter
}

// func NewRepositoryFilter(tab, name, value string) RepositoryFilter {
// 	return RepositoryFilter{Tab: tab, Name: name, Value: value}
// }

type RepositoryFilterAction int

// This is just a list of actions that can be performed on the repository filter
// there must be a better way to do this
const (
	AddAction RepositoryFilterAction = iota
	RemoveAction
	UpdateAction
	ClearAllAction
	CancelAction
)
