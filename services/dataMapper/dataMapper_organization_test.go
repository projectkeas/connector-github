package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestOrganizationMapping(t *testing.T) {
	src := map[string]interface{}{
		"login":              "octo-org",
		"id":                 6811672,
		"node_id":            "MDEyOk9yZ2FuaXphdGlvbjY4MTE2NzI=",
		"url":                "https://api.github.com/orgs/octo-org",
		"repos_url":          "https://api.github.com/orgs/octo-org/repos",
		"events_url":         "https://api.github.com/orgs/octo-org/events",
		"hooks_url":          "https://api.github.com/orgs/octo-org/hooks",
		"issues_url":         "https://api.github.com/orgs/octo-org/issues",
		"members_url":        "https://api.github.com/orgs/octo-org/members{/member}",
		"public_members_url": "https://api.github.com/orgs/octo-org/public_members{/member}",
		"avatar_url":         "https://avatars.githubusercontent.com/u/6811672?v=4",
		"description":        "Working better together!",
	}
	expected := map[string]interface{}{
		"avatarUrl":   "https://avatars.githubusercontent.com/u/6811672?v=4",
		"id":          6811672,
		"login":       "octo-org",
		"description": "Working better together!",
	}

	dest := mapOrganization(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
