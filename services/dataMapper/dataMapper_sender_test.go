package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestSenderMapping(t *testing.T) {
	src := map[string]interface{}{
		"login":               "Codertocat",
		"id":                  21031067,
		"node_id":             "MDQ6VXNlcjIxMDMxMDY3",
		"avatar_url":          "https://avatars1.githubusercontent.com/u/21031067?v=4",
		"gravatar_id":         "",
		"url":                 "https://api.github.com/users/Codertocat",
		"html_url":            "https://github.com/Codertocat",
		"followers_url":       "https://api.github.com/users/Codertocat/followers",
		"following_url":       "https://api.github.com/users/Codertocat/following{/other_user}",
		"gists_url":           "https://api.github.com/users/Codertocat/gists{/gist_id}",
		"starred_url":         "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
		"subscriptions_url":   "https://api.github.com/users/Codertocat/subscriptions",
		"organizations_url":   "https://api.github.com/users/Codertocat/orgs",
		"repos_url":           "https://api.github.com/users/Codertocat/repos",
		"events_url":          "https://api.github.com/users/Codertocat/events{/privacy}",
		"received_events_url": "https://api.github.com/users/Codertocat/received_events",
		"type":                "User",
		"site_admin":          false,
	}
	expected := map[string]interface{}{
		"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
		"id":          21031067,
		"isSiteAdmin": false,
		"login":       "Codertocat",
		"type":        "User",
		"url":         "https://github.com/Codertocat",
	}

	dest := mapSender(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}

func TestSenderMapping_WithNoAvatarUrl(t *testing.T) {
	src := map[string]interface{}{
		"login":               "Codertocat",
		"id":                  21031067,
		"node_id":             "MDQ6VXNlcjIxMDMxMDY3",
		"gravatar_id":         "",
		"url":                 "https://api.github.com/users/Codertocat",
		"html_url":            "https://github.com/Codertocat",
		"followers_url":       "https://api.github.com/users/Codertocat/followers",
		"following_url":       "https://api.github.com/users/Codertocat/following{/other_user}",
		"gists_url":           "https://api.github.com/users/Codertocat/gists{/gist_id}",
		"starred_url":         "https://api.github.com/users/Codertocat/starred{/owner}{/repo}",
		"subscriptions_url":   "https://api.github.com/users/Codertocat/subscriptions",
		"organizations_url":   "https://api.github.com/users/Codertocat/orgs",
		"repos_url":           "https://api.github.com/users/Codertocat/repos",
		"events_url":          "https://api.github.com/users/Codertocat/events{/privacy}",
		"received_events_url": "https://api.github.com/users/Codertocat/received_events",
		"type":                "User",
		"site_admin":          true,
	}
	expected := map[string]interface{}{
		"id":          21031067,
		"isSiteAdmin": true,
		"login":       "Codertocat",
		"type":        "User",
		"url":         "https://github.com/Codertocat",
	}

	dest := mapSender(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-expected +received):\n%s", diff)
	}
}
