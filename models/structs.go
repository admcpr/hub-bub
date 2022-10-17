package models

type Repository struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type User struct {
	Name  string
	Login string
}

type Organisation struct {
	Login        string `json:"login"`
	Url          string `json:"url"`
	Repositories []Repository
}
