package structs

import "time"

type FilterDate struct {
	Tab  string
	Name string
	From time.Time
	To   time.Time
}
