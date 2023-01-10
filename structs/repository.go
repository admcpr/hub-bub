package structs

import "time"

type Repository struct {
	Name                  string `title:"Name" description:"The name of the repository."`
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
		} `graphql:"branchProtectionRule"`
	} `graphql:"defaultBranchRef"`
	VulnerabilityAlerts struct {
		TotalCount int
	}
}
