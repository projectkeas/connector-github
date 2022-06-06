package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestAppMapping(t *testing.T) {
	src := map[string]interface{}{
		"id":      29310,
		"node_id": "MDM6QXBwMjkzMTA=",
		"owner": map[string]interface{}{
			"login":               "Octocoders",
			"id":                  38302899,
			"node_id":             "MDEyOk9yZ2FuaXphdGlvbjM4MzAyODk5",
			"avatar_url":          "https://avatars1.githubusercontent.com/u/38302899?v=4",
			"gravatar_id":         "",
			"url":                 "https://api.github.com/users/Octocoders",
			"html_url":            "https://github.com/Octocoders",
			"followers_url":       "https://api.github.com/users/Octocoders/followers",
			"following_url":       "https://api.github.com/users/Octocoders/following{/other_user}",
			"gists_url":           "https://api.github.com/users/Octocoders/gists{/gist_id}",
			"starred_url":         "https://api.github.com/users/Octocoders/starred{/owner}{/repo}",
			"subscriptions_url":   "https://api.github.com/users/Octocoders/subscriptions",
			"organizations_url":   "https://api.github.com/users/Octocoders/orgs",
			"repos_url":           "https://api.github.com/users/Octocoders/repos",
			"events_url":          "https://api.github.com/users/Octocoders/events{/privacy}",
			"received_events_url": "https://api.github.com/users/Octocoders/received_events",
			"type":                "Organization",
			"site_admin":          false,
		},
		"name":         "octocoders-linter",
		"description":  "",
		"external_url": "https://octocoders.github.io",
		"html_url":     "https://github.com/apps/octocoders-linter",
		"created_at":   "2019-04-19T19:36:24Z",
		"updated_at":   "2019-04-19T19:36:56Z",
		"permissions": map[string]interface{}{
			"administration":              "write",
			"checks":                      "write",
			"contents":                    "write",
			"deployments":                 "write",
			"issues":                      "write",
			"members":                     "write",
			"metadata":                    "read",
			"organization_administration": "write",
			"organization_hooks":          "write",
			"organization_plan":           "read",
			"organization_projects":       "write",
			"organization_user_blocking":  "write",
			"pages":                       "write",
			"pull_requests":               "write",
			"repository_hooks":            "write",
			"repository_projects":         "write",
			"statuses":                    "write",
			"team_discussions":            "write",
			"vulnerability_alerts":        "read",
		},
		"events": []string{},
	}
	expected := map[string]interface{}{
		"createdAt":   "2019-04-19T19:36:24Z",
		"description": "",
		"events":      []string{},
		"id":          29310,
		"name":        "octocoders-linter",
		"owner": map[string]any{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/38302899?v=4",
			"id":          38302899,
			"isSiteAdmin": false,
			"login":       "Octocoders",
			"type":        "Organization",
			"url":         "https://github.com/Octocoders",
		},
		"permissions": map[string]any{
			"administration":              "write",
			"checks":                      "write",
			"contents":                    "write",
			"deployments":                 "write",
			"issues":                      "write",
			"members":                     "write",
			"metadata":                    "read",
			"organization_administration": "write",
			"organization_hooks":          "write",
			"organization_plan":           "read",
			"organization_projects":       "write",
			"organization_user_blocking":  "write",
			"pages":                       "write",
			"pull_requests":               "write",
			"repository_hooks":            "write",
			"repository_projects":         "write",
			"statuses":                    "write",
			"team_discussions":            "write",
			"vulnerability_alerts":        "read",
		},
		"updatedAt": "2019-04-19T19:36:56Z",
		"url":       "https://github.com/apps/octocoders-linter",
	}

	dest := mapApp(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
