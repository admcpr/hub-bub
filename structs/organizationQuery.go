package structs

import (
	"time"
)

type OrganizationQuery struct {
	Organization struct {
		Id           string
		Repositories struct {
			Edges []struct {
				Node OrganisationRepositoryNodeQuery `graphql:"node"`
			} `graphql:"edges"`
		} `graphql:"repositories(first: $first)"`
	} `graphql:"organization(login: $login)"`
}

type OrganisationRepositoryNodeQuery struct {
	Name                  string
	Url                   string
	Id                    string
	AutoMergeAllowed      bool
	DeleteBranchOnMerge   bool
	RebaseMergeAllowed    bool
	HasDiscussionsEnabled bool
	HasIssuesEnabled      bool
	HasWikiEnabled        bool
	HasProjectsEnabled    bool
	IsArchived            bool
	IsDisabled            bool
	IsFork                bool
	IsLocked              bool
	IsMirror              bool
	IsPrivate             bool
	IsTemplate            bool
	StargazerCount        int
	SquashMergeAllowed    bool
	UpdatedAt             time.Time
	DefaultBranchRef      struct {
		Name                 string
		BranchProtectionRule struct {
			AllowsDeletions                bool
			AllowsForcePushes              bool
			DismissesStaleReviews          bool
			IsAdminEnforced                bool
			RequiredApprovingReviewCount   int
			RequiresApprovingReviews       bool
			RequiresCodeOwnerReviews       bool
			RequiresCommitSignatures       bool
			RequiresConversationResolution bool
			RequiresLinearHistory          bool
			RequiresStatusChecks           bool
			RequireReviewsFromCodeOwners   bool
			RequiresStrictStatusChecks     bool
		} `graphql:"branchProtectionRule"`
	} `graphql:"defaultBranchRef"`
	VulnerabilityAlerts struct {
		TotalCount int
	} `graphql:"vulnerabilityAlerts"`
}

func (oq *OrganizationQuery) GetRepositories() []Repository {
	var repositories []Repository
	for _, edge := range oq.Organization.Repositories.Edges {
		repositories = append(repositories, edge.Node.GetRepository())
	}
	return repositories
}

func (ornq OrganisationRepositoryNodeQuery) GetRepository() Repository {
	return Repository{
		Overview: RepositoryOverview{
			Loaded:                true,
			UpdatedAt:             ornq.UpdatedAt,
			Name:                  ornq.Name,
			Url:                   ornq.Url,
			Id:                    ornq.Id,
			HasDiscussionsEnabled: ornq.HasDiscussionsEnabled,
			HasIssuesEnabled:      ornq.HasIssuesEnabled,
			HasWikiEnabled:        ornq.HasWikiEnabled,
			HasProjectsEnabled:    ornq.HasProjectsEnabled,
			IsArchived:            ornq.IsArchived,
			IsDisabled:            ornq.IsDisabled,
			IsFork:                ornq.IsFork,
			IsLocked:              ornq.IsLocked,
			IsMirror:              ornq.IsMirror,
			IsPrivate:             ornq.IsPrivate,
			IsTemplate:            ornq.IsTemplate,
			StargazerCount:        ornq.StargazerCount,
		},
		DefaultBranch: RepositoryDefaultBranch{
			Loaded:                         true,
			Name:                           ornq.DefaultBranchRef.Name,
			AllowsDeletions:                ornq.DefaultBranchRef.BranchProtectionRule.AllowsDeletions,
			AllowsForcePushes:              ornq.DefaultBranchRef.BranchProtectionRule.AllowsForcePushes,
			DismissesStaleReviews:          ornq.DefaultBranchRef.BranchProtectionRule.DismissesStaleReviews,
			IsAdminEnforced:                ornq.DefaultBranchRef.BranchProtectionRule.IsAdminEnforced,
			RequiredApprovingReviewCount:   ornq.DefaultBranchRef.BranchProtectionRule.RequiredApprovingReviewCount,
			RequiresApprovingReviews:       ornq.DefaultBranchRef.BranchProtectionRule.RequiresApprovingReviews,
			RequiresCodeOwnerReviews:       ornq.DefaultBranchRef.BranchProtectionRule.RequiresCodeOwnerReviews,
			RequiresCommitSignatures:       ornq.DefaultBranchRef.BranchProtectionRule.RequiresCommitSignatures,
			RequiresConversationResolution: ornq.DefaultBranchRef.BranchProtectionRule.RequiresConversationResolution,
			RequiresLinearHistory:          ornq.DefaultBranchRef.BranchProtectionRule.RequiresLinearHistory,
			RequiresStatusChecks:           ornq.DefaultBranchRef.BranchProtectionRule.RequiresStatusChecks,
		},
	}
}
