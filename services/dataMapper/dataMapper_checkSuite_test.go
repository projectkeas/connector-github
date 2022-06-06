package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestCheckSuiteMapping(t *testing.T) {
	src := map[string]interface{}{
		"id":          118578147,
		"node_id":     "MDEwOkNoZWNrU3VpdGUxMTg1NzgxNDc=",
		"head_branch": "changes",
		"head_sha":    "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
		"status":      "completed",
		"conclusion":  "success",
		"url":         "https://api.github.com/repos/Codertocat/Hello-World/check-suites/118578147",
		"before":      "6113728f27ae82c7b1a177c8d03f9e96e0adf246",
		"after":       "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
		"pull_requests": []map[string]interface{}{
			{
				"url":    "https://api.github.com/repos/Codertocat/Hello-World/pulls/2",
				"id":     279147437,
				"number": 2,
				"head": map[string]interface{}{
					"ref": "changes",
					"sha": "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
					"repo": map[string]interface{}{
						"id":   186853002,
						"url":  "https://api.github.com/repos/Codertocat/Hello-World",
						"name": "Hello-World",
					},
				},
				"base": map[string]interface{}{
					"ref": "master",
					"sha": "f95f852bd8fca8fcc58a9a2d6c842781e32a215e",
					"repo": map[string]interface{}{
						"id":   186853002,
						"url":  "https://api.github.com/repos/Codertocat/Hello-World",
						"name": "Hello-World",
					},
				},
			},
		},
		"app": map[string]interface{}{
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
		},
		"created_at":              "2019-05-15T15:20:31Z",
		"updated_at":              "2019-05-15T15:21:14Z",
		"latest_check_runs_count": 1,
		"check_runs_url":          "https://api.github.com/repos/Codertocat/Hello-World/check-suites/118578147/check-runs",
		"head_commit": map[string]interface{}{
			"id":        "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
			"tree_id":   "31b122c26a97cf9af023e9ddab94a82c6e77b0ea",
			"message":   "Update README.md",
			"timestamp": "2019-05-15T15:20:30Z",
			"author": map[string]interface{}{
				"name":  "Codertocat",
				"email": "21031067+Codertocat@users.noreply.github.com",
			},
			"committer": map[string]interface{}{
				"name":  "Codertocat",
				"email": "21031067+Codertocat@users.noreply.github.com",
			},
		},
	}
	expected := map[string]interface{}{
		"after":      "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
		"before":     "6113728f27ae82c7b1a177c8d03f9e96e0adf246",
		"conclusion": "success",
		"createdAt":  "2019-05-15T15:20:31Z",
		"headBranch": "changes",
		"headCommit": map[string]interface{}{
			"author": map[string]interface{}{
				"email": "21031067+Codertocat@users.noreply.github.com",
				"name":  "Codertocat",
			},
			"committer": map[string]any{
				"email": "21031067+Codertocat@users.noreply.github.com",
				"name":  "Codertocat",
			},
			"id":        "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
			"message":   "Update README.md",
			"timestamp": "2019-05-15T15:20:30Z",
			"treeId":    "31b122c26a97cf9af023e9ddab94a82c6e77b0ea",
		},
		"headSha":   "ec26c3e57ca3a959ca5aad62de7213c562f8c821",
		"id":        118578147,
		"status":    "completed",
		"updatedAt": "2019-05-15T15:21:14Z",
		"app": map[string]interface{}{
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
		},
	}

	dest := mapCheckSuite(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
