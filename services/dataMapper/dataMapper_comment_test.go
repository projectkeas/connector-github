package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestIssueCommentMapping(t *testing.T) {
	src := map[string]interface{}{
		"url":       "https://api.github.com/repos/Codertocat/Hello-World/issues/comments/492700400",
		"html_url":  "https://github.com/Codertocat/Hello-World/issues/1#issuecomment-492700400",
		"issue_url": "https://api.github.com/repos/Codertocat/Hello-World/issues/1",
		"id":        492700400,
		"node_id":   "MDEyOklzc3VlQ29tbWVudDQ5MjcwMDQwMA==",
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
		"created_at":         "2019-05-15T15:20:21Z",
		"updated_at":         "2019-05-15T15:20:21Z",
		"author_association": "OWNER",
		"body":               "You are totally right! I'll get this fixed right away.",
	}
	expected := map[string]interface{}{
		"authorAssociation": "OWNER",
		"body":              "You are totally right! I'll get this fixed right away.",
		"createdAt":         "2019-05-15T15:20:21Z",
		"id":                492700400,
		"updatedAt":         "2019-05-15T15:20:21Z",
		"url":               "https://github.com/Codertocat/Hello-World/issues/1#issuecomment-492700400",
		"user": map[string]interface{}{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"id":          21031067,
			"isSiteAdmin": false,
			"login":       "Codertocat",
			"type":        "User",
			"url":         "https://github.com/Codertocat",
		},
	}

	dest := mapComment(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}

func TestCommitCommentMapping(t *testing.T) {
	src := map[string]interface{}{
		"url":      "https://api.github.com/repos/Codertocat/Hello-World/comments/33548674",
		"html_url": "https://github.com/Codertocat/Hello-World/commit/6113728f27ae82c7b1a177c8d03f9e96e0adf246#commitcomment-33548674",
		"id":       33548674,
		"node_id":  "MDEzOkNvbW1pdENvbW1lbnQzMzU0ODY3NA==",
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
		"position":           2,
		"line":               1,
		"path":               "test",
		"commit_id":          "6113728f27ae82c7b1a177c8d03f9e96e0adf246",
		"created_at":         "2019-05-15T15:20:39Z",
		"updated_at":         "2019-05-15T15:20:39Z",
		"author_association": "OWNER",
		"body":               "This is a really good change! :+1:",
	}
	expected := map[string]interface{}{
		"authorAssociation": "OWNER",
		"body":              "This is a really good change! :+1:",
		"commitId":          "6113728f27ae82c7b1a177c8d03f9e96e0adf246",
		"createdAt":         "2019-05-15T15:20:39Z",
		"id":                33548674,
		"line":              1,
		"path":              "test",
		"position":          2,
		"updatedAt":         "2019-05-15T15:20:39Z",
		"url":               "https://github.com/Codertocat/Hello-World/commit/6113728f27ae82c7b1a177c8d03f9e96e0adf246#commitcomment-33548674",
		"user": map[string]interface{}{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"id":          21031067,
			"isSiteAdmin": false,
			"login":       "Codertocat",
			"type":        "User",
			"url":         "https://github.com/Codertocat",
		},
	}

	dest := mapComment(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}

func TestDiscussionCommentMapping(t *testing.T) {
	src := map[string]interface{}{
		"id":                  544078,
		"node_id":             "MDE3OkRpc2N1c3Npb25Db21tZW50NTQ0MDc4",
		"html_url":            "https://github.com/octo-org/octo-repo/discussions/90#discussioncomment-544078",
		"parent_id":           "",
		"child_comment_count": 0,
		"repository_url":      "octo-org/octo-repo",
		"discussion_id":       3297442,
		"author_association":  "COLLABORATOR",
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
		"created_at": "2021-03-29T14:16:31Z",
		"updated_at": "2021-03-29T14:16:31Z",
		"body":       "I have so many questions to ask you!",
	}
	expected := map[string]interface{}{
		"authorAssociation": "COLLABORATOR",
		"body":              "I have so many questions to ask you!",
		"createdAt":         "2021-03-29T14:16:31Z",
		"discussionId":      3297442,
		"id":                544078,
		"updatedAt":         "2021-03-29T14:16:31Z",
		"url":               "https://github.com/octo-org/octo-repo/discussions/90#discussioncomment-544078",
		"user": map[string]interface{}{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"id":          21031067,
			"isSiteAdmin": false,
			"login":       "Codertocat",
			"type":        "User",
			"url":         "https://github.com/Codertocat",
		},
	}

	dest := mapComment(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}

func TestPullRequestReviewCommentMapping(t *testing.T) {
	src := map[string]interface{}{
		"url":                    "https://api.github.com/repos/Codertocat/Hello-World/pulls/comments/284312630",
		"pull_request_review_id": 237895671,
		"id":                     284312630,
		"node_id":                "MDI0OlB1bGxSZXF1ZXN0UmV2aWV3Q29tbWVudDI4NDMxMjYzMA==",
		"diff_hunk":              "@@ -1 +1 @@\n-# Hello-World",
		"path":                   "README.md",
		"position":               1,
		"original_position":      1,
		"commit_id":              "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
		"original_commit_id":     "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
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
		"body":               "Maybe you should use more emoji on this line.",
		"created_at":         "2019-05-15T15:20:37Z",
		"updated_at":         "2019-05-15T15:20:38Z",
		"html_url":           "https://github.com/Codertocat/Hello-World/pull/2#discussion_r284312630",
		"pull_request_url":   "https://api.github.com/repos/Codertocat/Hello-World/pulls/2",
		"author_association": "OWNER",
		"_links": map[string]interface{}{
			"self": map[string]interface{}{
				"href": "https://api.github.com/repos/Codertocat/Hello-World/pulls/comments/284312630",
			},
			"html": map[string]interface{}{
				"href": "https://github.com/Codertocat/Hello-World/pull/2#discussion_r284312630",
			},
			"pull_request": map[string]interface{}{
				"href": "https://api.github.com/repos/Codertocat/Hello-World/pulls/2",
			},
		},
	}
	expected := map[string]interface{}{
		"authorAssociation":   "OWNER",
		"body":                "Maybe you should use more emoji on this line.",
		"commitId":            "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
		"createdAt":           "2019-05-15T15:20:37Z",
		"diffHunk":            "@@ -1 +1 @@\n-# Hello-World",
		"id":                  284312630,
		"originalCommitId":    "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
		"originalPosition":    1,
		"path":                "README.md",
		"position":            1,
		"pullRequestReviewId": 237895671,
		"pullRequestUrl":      "https://api.github.com/repos/Codertocat/Hello-World/pulls/2",
		"updatedAt":           "2019-05-15T15:20:38Z",
		"url":                 "https://github.com/Codertocat/Hello-World/pull/2#discussion_r284312630",
		"user": map[string]interface{}{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"id":          21031067,
			"isSiteAdmin": false,
			"login":       "Codertocat",
			"type":        "User",
			"url":         "https://github.com/Codertocat",
		},
	}

	dest := mapComment(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
