package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepoSetting(t *testing.T) {
	name := "Name"
	title := "Title"

	repoSetting := NewRepoSetting(name, title)

	assert.Equal(t, name, repoSetting.Name)
	assert.Equal(t, title, repoSetting.Value)
}

func TestNewRepoSettingsTab(t *testing.T) {
	name := "Name"
	repositorySettings := []RepositorySetting{
		NewRepoSetting("name", "value"),
	}

	repoSettingsTab := NewRepoSettingsTab(name, repositorySettings)

	assert.Equal(t, name, repoSettingsTab.Name)
	assert.Equal(t, repositorySettings, repoSettingsTab.Settings)
}

// func TestBuildRepoSettings(t *testing.T) {
// 	ornq:= RepositoryQuery{
// 		Name: "Name",
// 		Url: "Url",
// 		Id: "Id",
// 		AutoMergeAllowed: true,
// 		DeleteBranchOnMerge: true,
// 		RebaseMergeAllowed: true,
// 		MergeCommitAllowed: true,
// 		HasDiscussionsEnabled: true,
// 		HasIssuesEnabled: true,
// 		HasWikiEnabled: true,
// 		HasProjectsEnabled: true,
// 		HasVulnerabilityAlertsEnabled: true,
// 		IsArchived: true,
// 		IsDisabled: true,
// 		IsFork: true,
// 		IsLocked: true,
// 		IsMirror: true,
// 		IsPrivate: true,
// 		IsTemplate: true,
// 		StargazerCount: 1,
// 		SquashMergeAllowed: true,
// 		UpdatedAt: "UpdatedAt",
// 		DefaultBranchRef: struct {
// 			Name                 string
// 			BranchProtectionRule BranchProtectionRuleQuery `graphql:"branchProtectionRule"`
// 		} {
// 			Name: "Name",
// 			BranchProtectionRule: struct {
// 				AllowsDeletions                bool
// 				AllowsForcePushes              bool
// 				DismissesStaleReviews          bool
// 				IsAdminEnforced                bool
// 				RequiredApprovingReviewCount   int
// 			} {
// 				AllowsDeletions: true,
// 				AllowsForcePushes: true,
// 				DismissesStaleReviews: true,
// 				IsAdminEnforced: true,
// 				RequiredApprovingReviewCount: 1,
// 			},
// 		},
// 		VulnerabilityAlerts: struct {
// 			TotalCount int
// 		} {
// 			TotalCount: 1,
// 		},
// 		PullRequests: struct {
// 			TotalCount int
// 		} {
// 			TotalCount: 1,
// 		},
// 	}

// 	type args struct {
// 		ornq RepositoryQuery
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []RepositorySettingsTab
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := BuildRepoSettings(tt.args.ornq); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("BuildRepoSettings() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_buildOverviewSettings(t *testing.T) {
// 	type args struct {
// 		ornq RepositoryQuery
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want RepositorySettingsTab
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := buildOverviewSettings(tt.args.ornq); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("buildOverviewSettings() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
