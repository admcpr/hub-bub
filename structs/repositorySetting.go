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

func buildRepositorySettings(ornq OrganisationRepositoryNodeQuery) []RepositorySetting {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Name", ornq.Name, "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Url", ornq.Url, "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Private?", utils.YesNo(ornq.IsPrivate), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Template?", utils.YesNo(ornq.IsTemplate), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Archived?", utils.YesNo(ornq.IsArchived), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Disabled?", utils.YesNo(ornq.IsDisabled), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Fork?", utils.YesNo(ornq.IsFork), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Last Updated?", fmt.Sprint(ornq.UpdatedAt), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Stars", fmt.Sprint(ornq.StargazerCount), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Wiki?", utils.YesNo(ornq.HasWikiEnabled), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Issues?", utils.YesNo(ornq.HasIssuesEnabled), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Projects?", utils.YesNo(ornq.HasProjectsEnabled), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Overview", "Discussions?", utils.YesNo(ornq.HasDiscussionsEnabled), "", "", true, false))

	return repositorySettings
}

func buildDefaultBranchSettings(ornq OrganisationRepositoryNodeQuery) []RepositorySetting {
	var repositorySettings []RepositorySetting

	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Name", ornq.DefaultBranchRef.Name, "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Protected?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.AllowsDeletions), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Allows Deletions?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.AllowsDeletions), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Allows Force Pushes?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.AllowsForcePushes), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Dismisses Stale Reviews?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.DismissesStaleReviews), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Admin Enforced?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.IsAdminEnforced), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Required Approving Review Count", fmt.Sprint(ornq.DefaultBranchRef.BranchProtectionRule.RequiredApprovingReviewCount), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Requires Approving Reviews?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.RequiresApprovingReviews), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Requires Code Owner Reviews?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.RequiresCodeOwnerReviews), "", "", true, false))
	repositorySettings = append(repositorySettings, NewRepositorySetting("Default Branch", "Requires Commit Signatures?", utils.YesNo(ornq.DefaultBranchRef.BranchProtectionRule.RequiresCommitSignatures), "", "", true, false))

	return repositorySettings
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
