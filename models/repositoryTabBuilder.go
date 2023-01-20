package models

import (
	"fmt"
	"strings"

	"github.com/admcpr/hub-bub/structs"
	"github.com/admcpr/hub-bub/utils"
)

func buildOverview(ornq structs.OrganisationRepositoryNodeQuery) string {
	stringbuilder := strings.Builder{}
	stringbuilder.WriteString(fmt.Sprintf("# %s", ornq.Name))
	stringbuilder.WriteString(fmt.Sprintf("| Wiki | %s |", utils.YesNo(ornq.HasWikiEnabled)))
	stringbuilder.WriteString(fmt.Sprintf("| Issues | %s |", utils.YesNo(ornq.HasIssuesEnabled)))
	stringbuilder.WriteString(fmt.Sprintf("| Projects | %s |", utils.YesNo(ornq.HasProjectsEnabled)))
	stringbuilder.WriteString(fmt.Sprintf("| Discussions | %s |", utils.YesNo(ornq.HasDiscussionsEnabled)))
	stringbuilder.WriteString(fmt.Sprintf("| Stars | %d |", ornq.StargazerCount))

	return stringbuilder.String()
}
