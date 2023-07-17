package structs

import "reflect"

type Filter interface {
	GetTab() string
	GetName() string
	GetType() reflect.Type
	Matches(setting Setting) bool
}
