package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestReviewMapping(t *testing.T) {
	src := map[string]interface{}{
		"id":      237895671,
		"node_id": "MDE3OlB1bGxSZXF1ZXN0UmV2aWV3MjM3ODk1Njcx",
		"user": map[string]interface{}{
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
		"body":               "body of the review",
		"commit_id":          "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
		"submitted_at":       "2019-05-15T15:20:38Z",
		"state":              "commented",
		"html_url":           "https://github.com/Codertocat/Hello-World/pull/2#pullrequestreview-237895671",
		"pull_request_url":   "https://api.github.com/repos/Codertocat/Hello-World/pulls/2",
		"author_association": "OWNER",
		"_links": map[string]interface{}{
			"html": map[string]interface{}{
				"href": "https://github.com/Codertocat/Hello-World/pull/2#pullrequestreview-237895671",
			},
			"pull_request": map[string]interface{}{
				"href": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2",
			},
		},
	}
	expected := map[string]interface{}{
		"authorAssociation": "OWNER",
		"body":              "body of the review",
		"commitId":          "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
		"id":                237895671,
		"pullRequestUrl":    "https://api.github.com/repos/Codertocat/Hello-World/pulls/2",
		"state":             "commented",
		"submittedAt":       "2019-05-15T15:20:38Z",
		"url":               "https://github.com/Codertocat/Hello-World/pull/2#pullrequestreview-237895671",
		"user": map[string]interface{}{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"id":          21031067,
			"isSiteAdmin": false,
			"login":       "Codertocat",
			"type":        "User",
			"url":         "https://github.com/Codertocat",
		},
	}

	dest := mapReview(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
