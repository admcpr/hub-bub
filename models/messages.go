package models

type OrgListMsg struct {
	Organisations []Organisation
}

type RepositoryListMsg struct {
	Repositories []Repository
}

type AuthenticationErrorMsg struct {
	Err error
}

type AuthenticationMsg struct {
	User User
}

type ErrMsg struct{ Err error }
