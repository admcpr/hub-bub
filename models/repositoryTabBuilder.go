package models

import (
	"fmt"
	"strings"

	"github.com/admcpr/hub-bub/structs"
	"github.com/admcpr/hub-bub/utils"
)

func buildOverview(ornq structs.OrganisationRepositoryNodeQuery) string {
	// TODO: Make this build a list
	stringbuilder := strings.Builder{}
	stringbuilder.WriteString(fmt.Sprintf("# %s", ornq.Name))
	stringbuilder.WriteString(fmt.Sprintf("| Wiki | %s |", utils.YesNo(ornq.HasWikiEnabled)))
	stringbuilder.WriteString(fmt.Sprintf("| Issues | %s |", utils.YesNo(ornq.HasIssuesEnabled)))
	stringbuilder.WriteString(fmt.Sprintf("| Projects | %s |", utils.YesNo(ornq.HasProjectsEnabled)))
	stringbuilder.WriteString(fmt.Sprintf("| Discussions | %s |", utils.YesNo(ornq.HasDiscussionsEnabled)))
	stringbuilder.WriteString(fmt.Sprintf("| Stars | %d |", ornq.StargazerCount))
	stringbuilder.WriteString(fmt.Sprintf("| Private? | %s |", utils.YesNo(ornq.IsPrivate)))
	stringbuilder.WriteString(fmt.Sprintf("| Fork? | %s |", utils.YesNo(ornq.IsFork)))
	stringbuilder.WriteString(fmt.Sprintf("| Archived? | %s |", utils.YesNo(ornq.IsArchived)))
	stringbuilder.WriteString(fmt.Sprintf("| Locked? | %s |", utils.YesNo(ornq.IsLocked)))
	stringbuilder.WriteString(fmt.Sprintf("| Template? | %s |", utils.YesNo(ornq.IsTemplate)))
	// stringbuilder.WriteString(fmt.Sprintf("| Description | %s |", ornq.Description))
	// stringbuilder.WriteString(fmt.Sprintf("| Created | %s |", ornq.CreatedAt))
	// stringbuilder.WriteString(fmt.Sprintf("| Updated | %s |", ornq.UpdatedAt))
	// stringbuilder.WriteString(fmt.Sprintf("| Pushed | %s |", ornq.PushedAt))
	// stringbuilder.WriteString(fmt.Sprintf("| Primary Language | %s |", ornq.PrimaryLanguage.Name))
	// stringbuilder.WriteString(fmt.Sprintf("| License | %s |", ornq.LicenseInfo.Name))
	// stringbuilder.WriteString(fmt.Sprintf("| URL | %s |", ornq.Url))
	// stringbuilder.WriteString(fmt.Sprintf("| Homepage URL | %s |", ornq.HomepageUrl))
	// stringbuilder.WriteString(fmt.Sprintf("| Owner | %s |", ornq.Owner.Login))
	// stringbuilder.WriteString(fmt.Sprintf("| Owner URL | %s |", ornq.Owner.Url))
	// stringbuilder.WriteString(fmt.Sprintf("| Owner Type | %s |", ornq.Owner.Type))
	// stringbuilder.WriteString(fmt.Sprintf("| Default Branch | %s |", ornq.DefaultBranchRef.Name))

	return stringbuilder.String()
}
