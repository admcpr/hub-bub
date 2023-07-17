package structs

import (
	"reflect"
	"time"
)

type FilterDate struct {
	Tab  string
	Name string
	From time.Time
	To   time.Time
}

func NewFilterDate(tab, name string, from, to time.Time) FilterDate {
	return FilterDate{Tab: tab, Name: name, From: from, To: to}
}

func (f FilterDate) GetTab() string {
	return f.Tab
}

func (f FilterDate) GetName() string {
	return f.Name
}

func (f FilterDate) GetType() reflect.Type {
	return reflect.TypeOf(f.From)
}

func (f FilterDate) Matches(setting Setting) bool {
	if setting.Type != reflect.TypeOf(f.From) {
		return false
	}

	date := setting.Value.(time.Time)

	return date.After(f.From) && date.Before(f.To)
}
