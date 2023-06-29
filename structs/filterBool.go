package structs

import "reflect"

type FilterBool struct {
	Tab   string
	Name  string
	Value bool
}

func NewFilterBool(tab, name string, value bool) FilterBool {
	return FilterBool{Tab: tab, Name: name, Value: value}
}

func (f FilterBool) GetTab() string {
	return f.Tab
}

func (f FilterBool) GetName() string {
	return f.Name
}

func (f FilterBool) GetType() reflect.Type {
	return reflect.TypeOf(f.Value)
}
