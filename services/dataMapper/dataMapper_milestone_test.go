package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestMilestoneMapping(t *testing.T) {
	src := map[string]interface{}{
		"url":         "https://api.github.com/repos/Codertocat/Hello-World/milestones/1",
		"html_url":    "https://github.com/Codertocat/Hello-World/milestone/1",
		"labels_url":  "https://api.github.com/repos/Codertocat/Hello-World/milestones/1/labels",
		"id":          4317517,
		"node_id":     "MDk6TWlsZXN0b25lNDMxNzUxNw==",
		"number":      1,
		"title":       "v1.0",
		"description": "Add new space flight simulator",
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
		"open_issues":   0,
		"closed_issues": 0,
		"state":         "open",
		"created_at":    "2019-05-15T15:20:17Z",
		"updated_at":    "2019-05-15T15:20:17Z",
		"due_on":        "2019-05-23T07:00:00Z",
		"closed_at":     "2019-05-23T07:00:00Z",
	}
	expected := map[string]interface{}{
		"closedIssues": 0,
		"closedAt":     "2019-05-23T07:00:00Z",
		"createdAt":    "2019-05-15T15:20:17Z",
		"creator": map[string]interface{}{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"id":          21031067,
			"isSiteAdmin": false,
			"login":       "Codertocat",
			"type":        "User",
			"url":         "https://github.com/Codertocat",
		},
		"description": "Add new space flight simulator",
		"dueOn":       "2019-05-23T07:00:00Z",
		"id":          4317517,
		"number":      1,
		"openIssues":  0,
		"state":       "open",
		"title":       "v1.0",
		"updatedAt":   "2019-05-15T15:20:17Z",
		"url":         "https://github.com/Codertocat/Hello-World/milestone/1",
	}

	dest := mapMilestone(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
