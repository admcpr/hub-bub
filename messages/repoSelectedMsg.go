package messages

import (
	"github.com/admcpr/hub-bub/structs"
)

type RepoSelectedMsg struct {
	RepositoryQuery structs.RepositoryQuery
}
