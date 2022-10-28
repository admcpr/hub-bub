package main

type User struct {
	Name  string
	Login string
}

type Repository struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Organisation struct {
	Login        string `json:"login"`
	Url          string `json:"url"`
	Repositories []Repository
}
