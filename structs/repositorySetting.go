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

func NewRepoSetting(name, value, url, propertyName string, loaded bool) RepositorySetting {
	return RepositorySetting{
		Name:         name,
		Value:        value,
		Url:          url,
		PropertyName: propertyName,
	}
}

func NewRepoSettingsTab(name string, settings []RepositorySetting) RepositorySettingsTab {
	return RepositorySettingsTab{
		Name:     name,
		Settings: settings,
	}
}

func BuildRepoSettings(ornq RepositoryQuery) []RepositorySettingsTab {
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
		NewRepoSetting("Private", YesNo(ornq.IsPrivate), "", "", true),
		NewRepoSetting("Template", YesNo(ornq.IsTemplate), "", "", true),
		NewRepoSetting("Archived", YesNo(ornq.IsArchived), "", "", true),
		NewRepoSetting("Disabled", YesNo(ornq.IsDisabled), "", "", true),
		NewRepoSetting("Fork", YesNo(ornq.IsFork), "", "", true),
		NewRepoSetting("Last updated", ornq.UpdatedAt.Format("2006/01/02"), "", "", true),
		NewRepoSetting("Stars", fmt.Sprint(ornq.StargazerCount), "", "", true),
		NewRepoSetting("Wiki", YesNo(ornq.HasWikiEnabled), "", "", true),
		NewRepoSetting("Issues", YesNo(ornq.HasIssuesEnabled), "", "", true),
		NewRepoSetting("Projects", YesNo(ornq.HasProjectsEnabled), "", "", true),
		NewRepoSetting("Discussions", YesNo(ornq.HasDiscussionsEnabled), "", "", true))

	return NewRepoSettingsTab("Overview", repositorySettings)
}

func buildPullRequestSettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepoSetting("Allow merge commits", YesNo(ornq.MergeCommitAllowed), "", "", true),
		NewRepoSetting("Allow squash merging", YesNo(ornq.SquashMergeAllowed), "", "", true),
		NewRepoSetting("Allow rebase merging", YesNo(ornq.RebaseMergeAllowed), "", "", true),
		NewRepoSetting("Allow auto-merge", YesNo(ornq.AutoMergeAllowed), "", "", true),
		NewRepoSetting("Automatically delete head branches", YesNo(ornq.DeleteBranchOnMerge), "", "", true),
		NewRepoSetting("Open pull requests", fmt.Sprint(ornq.PullRequests.TotalCount), "", "", true),
	)

	return NewRepoSettingsTab("Pull Requests", repositorySettings)
}

func buildDefaultBranchSettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	rule := ornq.DefaultBranchRef.BranchProtectionRule

	repositorySettings = append(repositorySettings,
		NewRepoSetting("Name", ornq.DefaultBranchRef.Name, "", "", true),
		NewRepoSetting("Require approving reviews", YesNo(rule.RequiresApprovingReviews), "", "", true),
		NewRepoSetting("Number of approvals required", fmt.Sprint(rule.RequiredApprovingReviewCount), "", "", true),
		NewRepoSetting("Dismiss stale requests", YesNo(rule.DismissesStaleReviews), "", "", true),
		NewRepoSetting("Require review from Code Owners", YesNo(rule.RequiresCodeOwnerReviews), "", "", true),
		NewRepoSetting("Restrict who can dismiss pull request reviews", YesNo(rule.RestrictsReviewDismissals), "", "", true),
		NewRepoSetting("Require approval of the most recent reviewable push", YesNo(rule.RequireLastPushApproval), "", "", true),
		// Allow specified actors to bypass required pull requests

		NewRepoSetting("Require status checks to pass before merging", YesNo(rule.RequiresStatusChecks), "", "", true),
		NewRepoSetting("Require conversation resolution before merging", YesNo(rule.RequiresConversationResolution), "", "", true),
		NewRepoSetting("Requires signed commits", YesNo(rule.RequiresCommitSignatures), "", "", true),
		NewRepoSetting("Require linear history", YesNo(rule.RequiresLinearHistory), "", "", true),
		NewRepoSetting("Require deployments to succeed before merging", YesNo(rule.RequiresDeployments), "", "", true),
		NewRepoSetting("Lock branch", YesNo(rule.LockBranch), "", "", true),
		NewRepoSetting("Do not allow bypassing the above settings", YesNo(rule.IsAdminEnforced), "", "", true),
		NewRepoSetting("Restrict who can push to matching branches", YesNo(rule.RestrictsPushes), "", "", true),

		NewRepoSetting("Allow force pushes", YesNo(rule.AllowsForcePushes), "", "", true),
		NewRepoSetting("Allow deletions", YesNo(rule.AllowsDeletions), "", "", true),
	)

	return NewRepoSettingsTab("Default Branch", repositorySettings)
}

func buildSecuritySettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepoSetting("Vulnerability alerts enabled", YesNo(ornq.HasVulnerabilityAlertsEnabled), "", "", true),
		NewRepoSetting("Vulnerability alert count", fmt.Sprint(ornq.VulnerabilityAlerts.TotalCount), "", "", true),
	)

	return NewRepoSettingsTab("Security", repositorySettings)
}
