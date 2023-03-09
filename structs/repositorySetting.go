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
		NewRepositorySetting("Last updated", fmt.Sprint(ornq.UpdatedAt), "", "", true),
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

	repositorySettings = append(repositorySettings,
		NewRepositorySetting("Name", ornq.DefaultBranchRef.Name, "", "", true),
		NewRepositorySetting("Require approving reviews", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.RequiresApprovingReviews), "", "", true),
		NewRepositorySetting("Required Approving Review Count", fmt.Sprint(ornq.DefaultBranchRef.BranchProtectionRule.RequiredApprovingReviewCount), "", "", true),
		NewRepositorySetting("Requires Code Owner Reviews?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.RequiresCodeOwnerReviews), "", "", true),
		NewRepositorySetting("Protected?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.AllowsDeletions), "", "", true),
		NewRepositorySetting("Dismisses Stale Reviews?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.DismissesStaleReviews), "", "", true),
		NewRepositorySetting("Admin Enforced?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.IsAdminEnforced), "", "", true),
		NewRepositorySetting("Requires Commit Signatures?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.RequiresCommitSignatures), "", "", true),
		NewRepositorySetting("Allow Force Pushes?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.AllowsForcePushes), "", "", true),
		NewRepositorySetting("Allow Deletions?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.AllowsDeletions), "", "", true))

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
