package structs

import (
	"fmt"

	"github.com/admcpr/hub-bub/utils"
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
		NewRepositorySetting("Private", utils.YesNo(ornq.IsPrivate), "", "", true),
		NewRepositorySetting("Template", utils.YesNo(ornq.IsTemplate), "", "", true),
		NewRepositorySetting("Archived", utils.YesNo(ornq.IsArchived), "", "", true),
		NewRepositorySetting("Disabled", utils.YesNo(ornq.IsDisabled), "", "", true),
		NewRepositorySetting("Fork", utils.YesNo(ornq.IsFork), "", "", true),
		NewRepositorySetting("Last updated", ornq.UpdatedAt.Format("2006/01/02"), "", "", true),
		NewRepositorySetting("Stars", fmt.Sprint(ornq.StargazerCount), "", "", true),
		NewRepositorySetting("Wiki", utils.YesNo(ornq.HasWikiEnabled), "", "", true),
		NewRepositorySetting("Issues", utils.YesNo(ornq.HasIssuesEnabled), "", "", true),
		NewRepositorySetting("Projects", utils.YesNo(ornq.HasProjectsEnabled), "", "", true),
		NewRepositorySetting("Discussions", utils.YesNo(ornq.HasDiscussionsEnabled), "", "", true))

	return NewRepositorySettingsTab("Overview", repositorySettings)
}

func buildPullRequestSettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepositorySetting("Allow merge commits", utils.YesNo(ornq.MergeCommitAllowed), "", "", true),
		NewRepositorySetting("Allow squash merging", utils.YesNo(ornq.SquashMergeAllowed), "", "", true),
		NewRepositorySetting("Allow rebase merging", utils.YesNo(ornq.RebaseMergeAllowed), "", "", true),
		NewRepositorySetting("Allow auto-merge", utils.YesNo(ornq.AutoMergeAllowed), "", "", true),
		NewRepositorySetting("Automatically delete head branches", utils.YesNo(ornq.DeleteBranchOnMerge), "", "", true),
	)

	return NewRepositorySettingsTab("Pull Requests", repositorySettings)
}

func buildDefaultBranchSettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	rule := ornq.DefaultBranchRef.BranchProtectionRule

	repositorySettings = append(repositorySettings,
		NewRepositorySetting("Name", ornq.DefaultBranchRef.Name, "", "", true),
		NewRepositorySetting("Require approving reviews", utils.YesNo(rule.RequiresApprovingReviews), "", "", true),
		NewRepositorySetting("Number of approvals required", fmt.Sprint(rule.RequiredApprovingReviewCount), "", "", true),
		NewRepositorySetting("Dismiss stale requests", utils.YesNo(rule.DismissesStaleReviews), "", "", true),
		NewRepositorySetting("Require review from Code Owners", utils.YesNo(rule.RequiresCodeOwnerReviews), "", "", true),
		// Restrict who can dismiss pull request reviews
		// Allow specified actors to bypass required pull requests
		// Require approval of the most recent reviewable push

		// Require status checks to pass before merging
		NewRepositorySetting("Require status checks to pass before merging", utils.YesNo(rule.RequiresStatusChecks), "", "", true),
		// Require conversation resolution before merging
		NewRepositorySetting("Requires signed commits", utils.YesNo(rule.RequiresCommitSignatures), "", "", true),
		NewRepositorySetting("Require linear history", utils.YesNo(rule.RequiresLinearHistory), "", "", true),
		// NewRepositorySetting("Require deployments to succeed before merging", utils.YesNo(rule.), "", "", true),
		// Require deployments to succeed before merging
		// Lock branch
		NewRepositorySetting("Do not allow bypassing the above settings", utils.YesNo(rule.IsAdminEnforced), "", "", true),
		// Restrict who can push to matching branches

		NewRepositorySetting("Allow force pushes", utils.YesNo(rule.AllowsForcePushes), "", "", true),
		NewRepositorySetting("Allow deletions", utils.YesNo(rule.AllowsDeletions), "", "", true),
	)

	return NewRepositorySettingsTab("Default Branch", repositorySettings)
}

func buildSecuritySettings(ornq RepositoryQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepositorySetting("Vulnerability Alerts Enabled?", utils.YesNo(ornq.HasVulnerabilityAlertsEnabled), "", "", true),
		NewRepositorySetting("Vulnerability Alert Count", fmt.Sprint(ornq.VulnerabilityAlerts.TotalCount), "", "", true),
	)

	return NewRepositorySettingsTab("Security", repositorySettings)
}
