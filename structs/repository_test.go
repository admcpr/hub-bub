package structs

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_RepositoryView(t *testing.T) {
	testRepo := Repository{
		// Basics: struct {
		// 	Loaded                bool,
		// 	UpdatedAt             time.Time,
		// 	Name                  string,
		// 	Url                   string,

		// Name:                  "test",
		// Url:                   "test",
		// Id:                    "test",
		// AutoMergeAllowed:      true,
		// DeleteBranchOnMerge:   true,
		// RebaseMergeAllowed:    true,
		// HasDiscussionsEnabled: true,
		// HasIssuesEnabled:      true,
		// HasWikiEnabled:        true,
		// HasProjectsEnabled:    true,
		// IsArchived:            true,
		// IsDisabled:            true,
		// IsFork:                true,
		// IsLocked:              true,
		// IsMirror:              true,
		// IsPrivate:             true,
		// IsTemplate:            true,
		// StargazerCount:        1,
		// SquashMergeAllowed:    true,
		// UpdatedAt:             time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	valueOfTestRepo := reflect.ValueOf(testRepo)
	typeOfTestRepo := valueOfTestRepo.Type()

	for i := 0; i < valueOfTestRepo.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\tTag: %s\n", typeOfTestRepo.Field(i).Name, valueOfTestRepo.Field(i).Interface(), string(typeOfTestRepo.Field(i).Tag))
	}

}
