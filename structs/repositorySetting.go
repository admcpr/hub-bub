package structs

import (
	"fmt"
)

type RepositorySetting struct {
	Name         string
	Value        string
	Url          string
	PropertyName string
}

type RepositorySettingsTab struct {
	Name     string
	Settings []RepositorySetting
}

func NewRepositorySetting(name, value, url, propertyName string, loaded bool) RepositorySetting {
	return RepositorySetting{
		Name:         name,
		Value:        value,
		Url:          url,
		PropertyName: propertyName,
	}
}

func NewRepositorySettingsTab(name string, settings []RepositorySetting) RepositorySettingsTab {
	return RepositorySettingsTab{
		Name:     name,
		Settings: settings,
	}
}

func BuildRepositorySettings(ornq RepositoryQuery) []RepositorySettingsTab {
	var respositorySettings []RepositorySettingsTab

	return append(respositorySettings,
		buildOverviewSettings(ornq),
		buildPullRequestSettings(ornq),
		buildDefaultBranchSettings(ornq),
		buildSecuritySettings(ornq))
}

func buildOverviewSettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepositorySetting("Private", YesNo(ornq.IsPrivate), "", "", true),
		NewRepositorySetting("Template", YesNo(ornq.IsTemplate), "", "", true),
		NewRepositorySetting("Archived", YesNo(ornq.IsArchived), "", "", true),
		NewRepositorySetting("Disabled", YesNo(ornq.IsDisabled), "", "", true),
		NewRepositorySetting("Fork", YesNo(ornq.IsFork), "", "", true),
		NewRepositorySetting("Last updated", ornq.UpdatedAt.Format("2006/01/02"), "", "", true),
		NewRepositorySetting("Stars", fmt.Sprint(ornq.StargazerCount), "", "", true),
		NewRepositorySetting("Wiki", YesNo(ornq.HasWikiEnabled), "", "", true),
		NewRepositorySetting("Issues", YesNo(ornq.HasIssuesEnabled), "", "", true),
		NewRepositorySetting("Projects", YesNo(ornq.HasProjectsEnabled), "", "", true),
		NewRepositorySetting("Discussions", YesNo(ornq.HasDiscussionsEnabled), "", "", true))

	return NewRepositorySettingsTab("Overview", repositorySettings)
}

func buildPullRequestSettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepositorySetting("Allow merge commits", YesNo(ornq.MergeCommitAllowed), "", "", true),
		NewRepositorySetting("Allow squash merging", YesNo(ornq.SquashMergeAllowed), "", "", true),
		NewRepositorySetting("Allow rebase merging", YesNo(ornq.RebaseMergeAllowed), "", "", true),
		NewRepositorySetting("Allow auto-merge", YesNo(ornq.AutoMergeAllowed), "", "", true),
		NewRepositorySetting("Automatically delete head branches", YesNo(ornq.DeleteBranchOnMerge), "", "", true),
		NewRepositorySetting("Open pull requests", fmt.Sprint(ornq.PullRequests.TotalCount), "", "", true),
	)

	return NewRepositorySettingsTab("Pull Requests", repositorySettings)
}

func buildDefaultBranchSettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	rule := ornq.DefaultBranchRef.BranchProtectionRule

	repositorySettings = append(repositorySettings,
		NewRepositorySetting("Name", ornq.DefaultBranchRef.Name, "", "", true),
		NewRepositorySetting("Require approving reviews", YesNo(rule.RequiresApprovingReviews), "", "", true),
		NewRepositorySetting("Number of approvals required", fmt.Sprint(rule.RequiredApprovingReviewCount), "", "", true),
		NewRepositorySetting("Dismiss stale requests", YesNo(rule.DismissesStaleReviews), "", "", true),
		NewRepositorySetting("Require review from Code Owners", YesNo(rule.RequiresCodeOwnerReviews), "", "", true),
		NewRepositorySetting("Restrict who can dismiss pull request reviews", YesNo(rule.RestrictsReviewDismissals), "", "", true),
		NewRepositorySetting("Require approval of the most recent reviewable push", YesNo(rule.RequireLastPushApproval), "", "", true),
		// Allow specified actors to bypass required pull requests

		NewRepositorySetting("Require status checks to pass before merging", YesNo(rule.RequiresStatusChecks), "", "", true),
		NewRepositorySetting("Require conversation resolution before merging", YesNo(rule.RequiresConversationResolution), "", "", true),
		NewRepositorySetting("Requires signed commits", YesNo(rule.RequiresCommitSignatures), "", "", true),
		NewRepositorySetting("Require linear history", YesNo(rule.RequiresLinearHistory), "", "", true),
		NewRepositorySetting("Require deployments to succeed before merging", YesNo(rule.RequiresDeployments), "", "", true),
		NewRepositorySetting("Lock branch", YesNo(rule.LockBranch), "", "", true),
		NewRepositorySetting("Do not allow bypassing the above settings", YesNo(rule.IsAdminEnforced), "", "", true),
		NewRepositorySetting("Restrict who can push to matching branches", YesNo(rule.RestrictsPushes), "", "", true),

		NewRepositorySetting("Allow force pushes", YesNo(rule.AllowsForcePushes), "", "", true),
		NewRepositorySetting("Allow deletions", YesNo(rule.AllowsDeletions), "", "", true),
	)

	return NewRepositorySettingsTab("Default Branch", repositorySettings)
}

func buildSecuritySettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepositorySetting("Vulnerability alerts enabled", YesNo(ornq.HasVulnerabilityAlertsEnabled), "", "", true),
		NewRepositorySetting("Vulnerability alert count", fmt.Sprint(ornq.VulnerabilityAlerts.TotalCount), "", "", true),
	)

	return NewRepositorySettingsTab("Security", repositorySettings)
}
