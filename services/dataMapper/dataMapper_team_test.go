package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestTeamMapping(t *testing.T) {
	src := map[string]interface{}{
		"name":             "github",
		"id":               3253328,
		"node_id":          "MDQ6VGVhbTMyNTMzMjg=",
		"slug":             "github",
		"description":      "Open-source team",
		"privacy":          "secret",
		"url":              "https://api.github.com/teams/3253328",
		"html_url":         "https://github.com/orgs/Octocoders/teams/github",
		"members_url":      "https://api.github.com/teams/3253328/members{/member}",
		"repositories_url": "https://api.github.com/teams/3253328/repos",
		"permission":       "pull",
	}
	expected := map[string]interface{}{
		"name":        "github",
		"id":          3253328,
		"slug":        "github",
		"description": "Open-source team",
		"privacy":     "secret",
		"url":         "https://github.com/orgs/Octocoders/teams/github",
		"permission":  "pull",
	}

	dest := mapTeam(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
