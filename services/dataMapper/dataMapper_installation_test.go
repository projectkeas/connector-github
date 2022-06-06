package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestInstallationMapping(t *testing.T) {
	src := map[string]interface{}{
		"id":      375706,
		"node_id": "MDIzOkludGVncmF0aW9uSW5zdGFsbGF0aW9uMzc1NzA2",
	}
	expected := map[string]interface{}{
		"id": 375706,
	}

	dest := mapInstallation(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}

func TestInstallationMapping_FullPayload(t *testing.T) {
	src := map[string]interface{}{
		"id": 2,
		"account": map[string]interface{}{
			"login":               "octocat",
			"id":                  1,
			"node_id":             "MDQ6VXNlcjE=",
			"avatar_url":          "https://github.com/images/error/octocat_happy.gif",
			"gravatar_id":         "",
			"url":                 "https://api.github.com/users/octocat",
			"html_url":            "https://github.com/octocat",
			"followers_url":       "https://api.github.com/users/octocat/followers",
			"following_url":       "https://api.github.com/users/octocat/following{/other_user}",
			"gists_url":           "https://api.github.com/users/octocat/gists{/gist_id}",
			"starred_url":         "https://api.github.com/users/octocat/starred{/owner}{/repo}",
			"subscriptions_url":   "https://api.github.com/users/octocat/subscriptions",
			"organizations_url":   "https://api.github.com/users/octocat/orgs",
			"repos_url":           "https://api.github.com/users/octocat/repos",
			"events_url":          "https://api.github.com/users/octocat/events{/privacy}",
			"received_events_url": "https://api.github.com/users/octocat/received_events",
			"type":                "User",
			"site_admin":          false,
		},
		"repository_selection": "selected",
		"access_tokens_url":    "https://api.github.com/installations/2/access_tokens",
		"repositories_url":     "https://api.github.com/installation/repositories",
		"html_url":             "https://github.com/settings/installations/2",
		"app_id":               5725,
		"target_id":            3880403,
		"target_type":          "User",
		"permissions": map[string]interface{}{
			"metadata": "read",
			"contents": "read",
			"issues":   "write",
		},
		"events": []string{
			"push",
			"pull_request",
		},
		"created_at":       1525109898,
		"updated_at":       1525109899,
		"single_file_name": "config.yml",
	}
	expected := map[string]interface{}{
		"id": 2,
		"account": map[string]interface{}{
			"id":          1,
			"login":       "octocat",
			"avatarUrl":   "https://github.com/images/error/octocat_happy.gif",
			"url":         "https://github.com/octocat",
			"type":        "User",
			"isSiteAdmin": false,
		},
		"repositorySelection": "selected",
		"appId":               5725,
		"targetId":            3880403,
		"targetType":          "User",
		"permissions": map[string]interface{}{
			"metadata": "read",
			"contents": "read",
			"issues":   "write",
		},
		"events": []string{
			"push",
			"pull_request",
		},
		"createdAt":      "2018-04-30T17:38:18Z",
		"updatedAt":      "2018-04-30T17:38:19Z",
		"singleFileName": "config.yml",
		"url":            "https://github.com/settings/installations/2",
	}

	dest := mapInstallation(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
