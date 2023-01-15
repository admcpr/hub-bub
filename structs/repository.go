package structs

import "time"

type Repository struct {
	Overview      RepositoryOverview
	DefaultBranch RepositoryDefaultBranch
	// AutoMergeAllowed    bool
	// DeleteBranchOnMerge bool
	// RebaseMergeAllowed  bool
	// SquashMergeAllowed bool
	// UpdatedAt          time.Time
	// DefaultBranchRef   struct {
	// 	BranchProtectionRule struct {
	// 	} `graphql:"branchProtectionRule"`
	// } `graphql:"defaultBranchRef"`
	// VulnerabilityAlerts struct {
	// 	TotalCount int
	// }
}

type RepositoryOverview struct {
	Loaded                bool
	UpdatedAt             time.Time
	Name                  string
	Url                   string
	Id                    string
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
}

type RepositoryDefaultBranch struct {
	Loaded                         bool
	Name                           string
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
}
