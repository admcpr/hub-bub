package main

type User struct {
	Name  string
	Login string
}

type Repository struct {
	Name                string `json:"name"`
	DefaultBranch       string `json:"default_branch"`
	DeleteBranchOnMerge bool   `json:"delete_branch_on_merge"`
	HasPages            bool   `json:"has_pages"`
	HasIssues           bool   `json:"has_issues"`
	HasProjects         bool   `json:"has_projects"`
	HasWiki             bool   `json:"has_wiki"`
	IsPrivate           bool   `json:"private"`
	IsTemplate          bool   `json:"is_template"`
	IsArchived          bool   `json:"archived"`
	AllowAutoMerge      bool   `json:"allow_auto_merge"`
	AllowRebaseMerge    bool   `json:"allow_rebase_merge"`
	AllowMergeCommit    bool   `json:"allow_merge_commit"`
	Url                 string `json:"url"`
}

type Organisation struct {
	Login        string `json:"login"`
	Url          string `json:"url"`
	Repositories []Repository
}
