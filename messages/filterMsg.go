package messages

import (
	"hub-bub/structs"
	"reflect"
)

type FilterMsg struct {
	Filter interface{}
	Action structs.RepositoryFilterAction
	Type   reflect.Type
}
