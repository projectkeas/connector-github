package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestDeploymentStatusMapping(t *testing.T) {
	src := map[string]interface{}{
		"url":     "https://api.github.com/repos/Codertocat/Hello-World/deployments/145988746/statuses/209916254",
		"id":      209916254,
		"node_id": "MDE2OkRlcGxveW1lbnRTdGF0dXMyMDk5MTYyNTQ=",
		"state":   "success",
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
		"description":    "test",
		"environment":    "production",
		"target_url":     "",
		"created_at":     "2019-05-15T15:20:55Z",
		"updated_at":     "2019-05-15T15:20:55Z",
		"deployment_url": "https://api.github.com/repos/Codertocat/Hello-World/deployments/145988746",
		"repository_url": "https://api.github.com/repos/Codertocat/Hello-World",
	}
	expected := map[string]interface{}{
		"createdAt": "2019-05-15T15:20:55Z",
		"creator": map[string]interface{}{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"id":          21031067,
			"isSiteAdmin": false,
			"login":       "Codertocat",
			"type":        "User",
			"url":         "https://github.com/Codertocat",
		},
		"deploymentUrl": "https://api.github.com/repos/Codertocat/Hello-World/deployments/145988746",
		"description":   "test",
		"environment":   "production",
		"id":            209916254,
		"repositoryUrl": "https://api.github.com/repos/Codertocat/Hello-World",
		"state":         "success",
		"targetUrl":     "",
		"updatedAt":     "2019-05-15T15:20:55Z",
	}

	dest := mapDeploymentStatus(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
