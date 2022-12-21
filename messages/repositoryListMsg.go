package messages

import "github.com/admcpr/hub-bub/queries"

type RepositoryListMsg struct {
	OrganizationQuery queries.OrganizationQuery
}
