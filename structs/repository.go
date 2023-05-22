package structs

import (
	"fmt"
	"time"
)

type SettingValue interface {
	bool | string | time.Time | int
}

type SettingGetter interface {
	GetName() string
	GetValue() string
}

type Setting[T SettingValue] struct {
	Name  string
	Value T
}

func NewSetting[T SettingValue](name string, value T) Setting[T] {
	return Setting[T]{Name: name, Value: value}
}

func (s Setting[T]) GetValue() string {
	formattedValue := ""

	switch value := any(s.Value).(type) {
	case bool:
		formattedValue = YesNo(value)
	case string:
		formattedValue = value
	case time.Time:
		formattedValue = value.Format("2006/01/02")
	case int:
		formattedValue = fmt.Sprint(value)
	}

	return formattedValue
}

func (s Setting[T]) GetName() string {
	return s.Name
}

type Repository struct {
	Name     string
	Settings map[string][]SettingGetter
}

func NewRepository(rq RepositoryQuery) Repository {
	rule := rq.DefaultBranchRef.BranchProtectionRule

	return Repository{
		Name: rq.Name,
		Settings: map[string][]SettingGetter{
			"Overview": {
				NewSetting("Private", rq.IsPrivate),
				NewSetting("Template", rq.IsTemplate),
				NewSetting("Archived", rq.IsArchived),
				NewSetting("Disabled", rq.IsDisabled),
				NewSetting("Fork", rq.IsFork),
				NewSetting("Last updated", rq.UpdatedAt),
				NewSetting("Stars", rq.StargazerCount),
				NewSetting("Wiki", rq.HasWikiEnabled),
				NewSetting("Issues", rq.HasIssuesEnabled),
				NewSetting("Projects", rq.HasProjectsEnabled),
				NewSetting("Discussions", rq.HasDiscussionsEnabled),
			},
			"Pull Requests": {
				NewSetting("Allow merge commits", rq.MergeCommitAllowed),
				NewSetting("Allow squash merging", rq.SquashMergeAllowed),
				NewSetting("Allow rebase merging", rq.RebaseMergeAllowed),
				NewSetting("Allow auto-merge", rq.AutoMergeAllowed),
				NewSetting("Automatically delete head branches", rq.DeleteBranchOnMerge),
				NewSetting("Open pull requests", rq.PullRequests.TotalCount),
			},
			"Default Branch": {
				NewSetting("Name", rq.DefaultBranchRef.Name),
				NewSetting("Require approving reviews", rule.RequiresApprovingReviews),
				NewSetting("Number of approvals required", rule.RequiredApprovingReviewCount),
				NewSetting("Dismiss stale requests", rule.DismissesStaleReviews),
				NewSetting("Require review from Code Owners", rule.RequiresCodeOwnerReviews),
				NewSetting("Restrict who can dismiss pull request reviews", rule.RestrictsReviewDismissals),
				NewSetting("Require approval of the most recent reviewable push", rule.RequireLastPushApproval),
				NewSetting("Require status checks to pass before merging", rule.RequiresStatusChecks),
				NewSetting("Require conversation resolution before merging", rule.RequiresConversationResolution),
				NewSetting("Requires signed commits", rule.RequiresCommitSignatures),
				NewSetting("Require linear history", rule.RequiresLinearHistory),
				NewSetting("Require deployments to succeed before merging", rule.RequiresDeployments),
				NewSetting("Lock branch", rule.LockBranch),
				NewSetting("Do not allow bypassing the above settings", rule.IsAdminEnforced),
				NewSetting("Restrict who can push to matching branches", rule.RestrictsPushes),
				NewSetting("Allow force pushes", rule.AllowsForcePushes),
				NewSetting("Allow deletions", rule.AllowsDeletions),
			},
			"Security": {
				NewSetting("Vulnerability alerts enabled", rq.HasVulnerabilityAlertsEnabled),
				NewSetting("Vulnerability alert count", rq.VulnerabilityAlerts.TotalCount),
			},
		},
	}
}

func (r Repository) GetTabs() []string {
	var tabs []string

	for tab := range r.Settings {
		tabs = append(tabs, tab)
	}

	return tabs
}

type Organization struct {
	Name         string
	Id           string
	Login        string
	Url          string
	Repositories []Repository
}
