package structs

type RepositoryFilter struct {
	Tab   string
	Name  string
	Value string
}

func NewRepositoryFilter(tab, name, value string) RepositoryFilter {
	return RepositoryFilter{Tab: tab, Name: name, Value: value}
}

type RepositoryFilterAction int

const (
	Add RepositoryFilterAction = iota
	Remove
	Update
	ClearAll
)
