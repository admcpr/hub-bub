package structs

import (
	"fmt"

	"github.com/admcpr/hub-bub/utils"
)

type RepositorySetting struct {
	Tab          string
	Name         string
	Value        string
	Url          string
	Loaded       bool
	FromRest     bool
	PropertyName string
}

type RepositorySettingsTab struct {
	Name     string
	Settings []RepositorySetting
}

func NewRepositorySetting(tab, name, value, url, propertyName string, loaded, fromRest bool) RepositorySetting {
	return RepositorySetting{
		Tab:          tab,
		Name:         name,
		Value:        value,
		Url:          url,
		Loaded:       loaded,
		FromRest:     fromRest,
		PropertyName: propertyName,
	}
}

func NewRepositorySettingsTab(name string, settings []RepositorySetting) RepositorySettingsTab {
	return RepositorySettingsTab{
		Name:     name,
		Settings: settings,
	}
}

func BuildRepositorySettings(ornq OrganisationRepositoryNodeQuery) []RepositorySettingsTab {
	var respositorySettings []RepositorySettingsTab

	return append(respositorySettings,
		buildOverviewSettings(ornq),
		buildDefaultBranchSettings(ornq),
		buildSecuritySettings(ornq))
}

func buildOverviewSettings(ornq OrganisationRepositoryNodeQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings,
		NewRepositorySetting("Overview", "Name", ornq.Name, "", "", true, false),
		NewRepositorySetting("Overview", "Url", ornq.Url, "", "", true, false),
		NewRepositorySetting("Overview", "Private?", utils.YesNo(ornq.IsPrivate), "", "", true, false),
		NewRepositorySetting("Overview", "Template?", utils.YesNo(ornq.IsTemplate), "", "", true, false),
		NewRepositorySetting("Overview", "Archived?", utils.YesNo(ornq.IsArchived), "", "", true, false),
		NewRepositorySetting("Overview", "Disabled?", utils.YesNo(ornq.IsDisabled), "", "", true, false),
		NewRepositorySetting("Overview", "Fork?", utils.YesNo(ornq.IsFork), "", "", true, false),
		NewRepositorySetting("Overview", "Last Updated?", fmt.Sprint(ornq.UpdatedAt), "", "", true, false),
		NewRepositorySetting("Overview", "Stars", fmt.Sprint(ornq.StargazerCount), "", "", true, false),
		NewRepositorySetting("Overview", "Wiki?", utils.YesNo(ornq.HasWikiEnabled), "", "", true, false),
		NewRepositorySetting("Overview", "Issues?", utils.YesNo(ornq.HasIssuesEnabled), "", "", true, false),
		NewRepositorySetting("Overview", "Projects?", utils.YesNo(ornq.HasProjectsEnabled), "", "", true, false),
		NewRepositorySetting("Overview", "Discussions?", utils.YesNo(ornq.HasDiscussionsEnabled), "", "", true, false))

	return NewRepositorySettingsTab("Overview", repositorySettings)
}

func buildDefaultBranchSettings(ornq OrganisationRepositoryNodeQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Name", ornq.DefaultBranchRef.Name, "", "", true, false))

	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Requires Approving Reviews?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.RequiresApprovingReviews), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Required Approving Review Count", fmt.Sprint(ornq.DefaultBranchRef.BranchProtectionRule.RequiredApprovingReviewCount), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Requires Code Owner Reviews?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.RequiresCodeOwnerReviews), "", "", true, false))

	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Protected?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.AllowsDeletions), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Dismisses Stale Reviews?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.DismissesStaleReviews), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Admin Enforced?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.IsAdminEnforced), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Requires Commit Signatures?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.RequiresCommitSignatures), "", "", true, false))

	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Allow Force Pushes?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.AllowsForcePushes), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Allow Deletions?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.AllowsDeletions), "", "", true, false))

	return NewRepositorySettingsTab("Default Branch", repositorySettings)
}

func buildSecuritySettings(ornq OrganisationRepositoryNodeQuery) RepositorySettingsTab {
	var repositorySettings []RepositorySetting

	return NewRepositorySettingsTab("Security", repositorySettings)
}

//  Name                  string
// 	Url                   string
// 	Id                    string
// 	AutoMergeAllowed      bool
// 	DeleteBranchOnMerge   bool
// 	RebaseMergeAllowed    bool
// 	- HasDiscussionsEnabled bool
// 	- HasIssuesEnabled      bool
// 	- HasWikiEnabled        bool
// 	- HasProjectsEnabled    bool
// 	- IsArchived            bool
// 	- IsDisabled            bool
// 	- IsFork                bool
// 	IsLocked              bool
// 	IsMirror              bool
// 	- IsPrivate             bool
// 	- IsTemplate            bool
// 	- StargazerCount        int
// 	SquashMergeAllowed    bool
// 	UpdatedAt             time.Time
// 	DefaultBranchRef      struct {
// 		Name                 string
// 		BranchProtectionRule struct {
// 			AllowsDeletions                bool
// 			AllowsForcePushes              bool
// 			DismissesStaleReviews          bool
// 			IsAdminEnforced                bool
// 			RequiredApprovingReviewCount   int
// 			RequiresApprovingReviews       bool
// 			RequiresCodeOwnerReviews       bool
// 			RequiresCommitSignatures       bool
// 			RequiresConversationResolution bool
// 			RequiresLinearHistory          bool
// 			RequiresStatusChecks           bool
// 		} `graphql:"branchProtectionRule"`
// 	} `graphql:"defaultBranchRef"`
// 	VulnerabilityAlerts struct {
// 		TotalCount int
// 	} `graphql:"vulnerabilityAlerts"`
