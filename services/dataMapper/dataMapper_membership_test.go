package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestMembershipMapping(t *testing.T) {
	src := map[string]interface{}{
		"url":              "https://api.github.com/orgs/Octocoders/memberships/hacktocat",
		"state":            "pending",
		"role":             "member",
		"organization_url": "https://api.github.com/orgs/Octocoders",
		"user": map[string]interface{}{
			"login":               "hacktocat",
			"id":                  39652351,
			"node_id":             "MDQ6VXNlcjM5NjUyMzUx",
			"avatar_url":          "https://avatars2.githubusercontent.com/u/39652351?v=4",
			"gravatar_id":         "",
			"url":                 "https://api.github.com/users/hacktocat",
			"html_url":            "https://github.com/hacktocat",
			"followers_url":       "https://api.github.com/users/hacktocat/followers",
			"following_url":       "https://api.github.com/users/hacktocat/following{/other_user}",
			"gists_url":           "https://api.github.com/users/hacktocat/gists{/gist_id}",
			"starred_url":         "https://api.github.com/users/hacktocat/starred{/owner}{/repo}",
			"subscriptions_url":   "https://api.github.com/users/hacktocat/subscriptions",
			"organizations_url":   "https://api.github.com/users/hacktocat/orgs",
			"repos_url":           "https://api.github.com/users/hacktocat/repos",
			"events_url":          "https://api.github.com/users/hacktocat/events{/privacy}",
			"received_events_url": "https://api.github.com/users/hacktocat/received_events",
			"type":                "User",
			"site_admin":          false,
		},
	}
	expected := map[string]interface{}{
		"role":  "member",
		"state": "pending",
		"user": map[string]interface{}{
			"avatarUrl":   "https://avatars2.githubusercontent.com/u/39652351?v=4",
			"id":          39652351,
			"isSiteAdmin": false,
			"login":       "hacktocat",
			"type":        "User",
			"url":         "https://github.com/hacktocat",
		},
	}

	dest := mapMembership(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
