package main

type RepositoryListMsg struct {
	Repositories []Repository
}

type AuthenticationMsg struct {
	User User
}

type ErrMsg struct{ Err error }

type OrgListMsg struct {
	Organisations []Organisation
}

type AuthenticationErrorMsg struct {
	Err error
}
