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

func NewRepoSetting(name, value string) RepositorySetting {
	return RepositorySetting{
		Name:  name,
		Value: value,
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
		NewRepoSetting("Private", YesNo(ornq.IsPrivate)),
		NewRepoSetting("Template", YesNo(ornq.IsTemplate)),
		NewRepoSetting("Archived", YesNo(ornq.IsArchived)),
		NewRepoSetting("Disabled", YesNo(ornq.IsDisabled)),
		NewRepoSetting("Fork", YesNo(ornq.IsFork)),
		NewRepoSetting("Last updated", ornq.UpdatedAt.Format("2006/01/02")),
		NewRepoSetting("Stars", fmt.Sprint(ornq.StargazerCount)),
		NewRepoSetting("Wiki", YesNo(ornq.HasWikiEnabled)),
		NewRepoSetting("Issues", YesNo(ornq.HasIssuesEnabled)),
		NewRepoSetting("Projects", YesNo(ornq.HasProjectsEnabled)),
		NewRepoSetting("Discussions", YesNo(ornq.HasDiscussionsEnabled)))

	return NewRepoSettingsTab("Overview", repositorySettings)
}

func buildPullRequestSettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepoSetting("Allow merge commits", YesNo(ornq.MergeCommitAllowed)),
		NewRepoSetting("Allow squash merging", YesNo(ornq.SquashMergeAllowed)),
		NewRepoSetting("Allow rebase merging", YesNo(ornq.RebaseMergeAllowed)),
		NewRepoSetting("Allow auto-merge", YesNo(ornq.AutoMergeAllowed)),
		NewRepoSetting("Automatically delete head branches", YesNo(ornq.DeleteBranchOnMerge)),
		NewRepoSetting("Open pull requests", fmt.Sprint(ornq.PullRequests.TotalCount)),
	)

	return NewRepoSettingsTab("Pull Requests", repositorySettings)
}

func buildDefaultBranchSettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	rule := ornq.DefaultBranchRef.BranchProtectionRule

	repositorySettings = append(repositorySettings,
		NewRepoSetting("Name", ornq.DefaultBranchRef.Name),
		NewRepoSetting("Require approving reviews", YesNo(rule.RequiresApprovingReviews)),
		NewRepoSetting("Number of approvals required", fmt.Sprint(rule.RequiredApprovingReviewCount)),
		NewRepoSetting("Dismiss stale requests", YesNo(rule.DismissesStaleReviews)),
		NewRepoSetting("Require review from Code Owners", YesNo(rule.RequiresCodeOwnerReviews)),
		NewRepoSetting("Restrict who can dismiss pull request reviews", YesNo(rule.RestrictsReviewDismissals)),
		NewRepoSetting("Require approval of the most recent reviewable push", YesNo(rule.RequireLastPushApproval)),
		// Allow specified actors to bypass required pull requests

		NewRepoSetting("Require status checks to pass before merging", YesNo(rule.RequiresStatusChecks)),
		NewRepoSetting("Require conversation resolution before merging", YesNo(rule.RequiresConversationResolution)),
		NewRepoSetting("Requires signed commits", YesNo(rule.RequiresCommitSignatures)),
		NewRepoSetting("Require linear history", YesNo(rule.RequiresLinearHistory)),
		NewRepoSetting("Require deployments to succeed before merging", YesNo(rule.RequiresDeployments)),
		NewRepoSetting("Lock branch", YesNo(rule.LockBranch)),
		NewRepoSetting("Do not allow bypassing the above settings", YesNo(rule.IsAdminEnforced)),
		NewRepoSetting("Restrict who can push to matching branches", YesNo(rule.RestrictsPushes)),

		NewRepoSetting("Allow force pushes", YesNo(rule.AllowsForcePushes)),
		NewRepoSetting("Allow deletions", YesNo(rule.AllowsDeletions)),
	)

	return NewRepoSettingsTab("Default Branch", repositorySettings)
}

func buildSecuritySettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepoSetting("Vulnerability alerts enabled", YesNo(ornq.HasVulnerabilityAlertsEnabled)),
		NewRepoSetting("Vulnerability alert count", fmt.Sprint(ornq.VulnerabilityAlerts.TotalCount)),
	)

	return NewRepoSettingsTab("Security", repositorySettings)
}
