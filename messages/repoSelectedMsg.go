package messages

import (
	"hub-bub/structs"
)

type RepoSelectMsg struct {
	Repository structs.Repository
	Width      int
	Height     int
}
