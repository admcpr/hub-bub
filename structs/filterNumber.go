package structs

import "reflect"

type FilterNumber struct {
	Tab  string
	Name string
	From int
	To   int
}

func NewFilterNumber(tab, name string, from, to int) FilterNumber {
	return FilterNumber{Tab: tab, Name: name, From: from, To: to}
}

func (f FilterNumber) GetTab() string {
	return f.Tab
}

func (f FilterNumber) GetName() string {
	return f.Name
}

func (f FilterNumber) GetType() reflect.Type {
	return reflect.TypeOf(f.From)
}
