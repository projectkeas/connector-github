package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestDeploymentMapping(t *testing.T) {
	src := map[string]interface{}{
		"url":                  "https://api.github.com/repos/Codertocat/Hello-World/deployments/145988746",
		"id":                   145988746,
		"node_id":              "MDEwOkRlcGxveW1lbnQxNDU5ODg3NDY=",
		"sha":                  "f95f852bd8fca8fcc58a9a2d6c842781e32a215e",
		"ref":                  "master",
		"task":                 "deploy",
		"payload":              map[string]interface{}{},
		"original_environment": "production",
		"environment":          "production",
		"description":          "description",
		"creator": map[string]interface{}{
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
		},
		"created_at":     "2019-05-15T15:20:53Z",
		"updated_at":     "2019-05-15T15:20:53Z",
		"statuses_url":   "https://api.github.com/repos/Codertocat/Hello-World/deployments/145988746/statuses",
		"repository_url": "https://api.github.com/repos/Codertocat/Hello-World",
	}
	expected := map[string]interface{}{
		"createdAt": "2019-05-15T15:20:53Z",
		"creator": map[string]interface{}{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"id":          21031067,
			"isSiteAdmin": false,
			"login":       "Codertocat",
			"type":        "User",
			"url":         "https://github.com/Codertocat",
		},
		"description":         "description",
		"environment":         "production",
		"id":                  145988746,
		"originalEnvironment": "production",
		"payload":             map[string]interface{}{},
		"ref":                 "master",
		"sha":                 "f95f852bd8fca8fcc58a9a2d6c842781e32a215e",
		"task":                "deploy",
		"updatedAt":           "2019-05-15T15:20:53Z",
	}

	dest := mapDeployment(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
