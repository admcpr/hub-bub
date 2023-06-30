package structs

import "reflect"

type FilterInt struct {
	Tab  string
	Name string
	From int
	To   int
}

func NewFilterInt(tab, name string, from, to int) FilterInt {
	return FilterInt{Tab: tab, Name: name, From: from, To: to}
}

func (f FilterInt) GetTab() string {
	return f.Tab
}

func (f FilterInt) GetName() string {
	return f.Name
}

func (f FilterInt) GetType() reflect.Type {
	return reflect.TypeOf(f.From)
}
