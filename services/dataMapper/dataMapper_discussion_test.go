package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestDiscussionMapping(t *testing.T) {
	src := map[string]interface{}{
		"repository_url": "https://api.github.com/repos/octo-org/octo-repo",
		"category": map[string]interface{}{
			"id":            32784361,
			"node_id":       "DIC_kwAqGA",
			"repository_id": 17273051,
			"emoji":         ":speech_balloon:",
			"name":          "General",
			"description":   "Chat about anything and everything here",
			"created_at":    "2021-03-24T12:41:54.000-05:00",
			"updated_at":    "2021-03-24T12:41:54.000-05:00",
			"slug":          "general",
			"is_answerable": false,
		},
		"answer_html_url":  "none1",
		"answer_chosen_at": "none2",
		"answer_chosen_by": "none3",
		"html_url":         "https://github.com/octo-org/octo-repo/discussions/90",
		"id":               3297442,
		"node_id":          "MDEwOkRpc2N1c3Npb24zMjk3NDQy",
		"number":           90,
		"title":            "Welcome to discussions!",
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
		"state":              "open",
		"locked":             false,
		"comments":           0,
		"created_at":         "2021-03-29T14:16:08Z",
		"updated_at":         "2021-03-29T14:16:08Z",
		"author_association": "COLLABORATOR",
		"active_lock_reason": "none",
		"body":               "We're glad to have you here!",
	}
	expected := map[string]interface{}{
		"activeLockReason":  "none",
		"answerChosenAt":    "none2",
		"answerChosenBy":    "none3",
		"answerUrl":         "none1",
		"authorAssociation": "COLLABORATOR",
		"body":              "We're glad to have you here!",
		"category": map[string]interface{}{
			"createdAt":    "2021-03-24T12:41:54.000-05:00",
			"description":  "Chat about anything and everything here",
			"emoji":        ":speech_balloon:",
			"id":           32784361,
			"isAnswerable": false,
			"name":         "General",
			"slug":         "general",
			"updatedAt":    "2021-03-24T12:41:54.000-05:00",
		},
		"comments":  0,
		"createdAt": "2021-03-29T14:16:08Z",
		"id":        3297442,
		"isLocked":  false,
		"number":    90,
		"state":     "open",
		"title":     "Welcome to discussions!",
		"updatedAt": "2021-03-29T14:16:08Z",
		"url":       "https://github.com/octo-org/octo-repo/discussions/90",
		"user": map[string]interface{}{
			"avatarUrl":   "https://avatars1.githubusercontent.com/u/21031067?v=4",
			"id":          21031067,
			"isSiteAdmin": false,
			"login":       "Codertocat",
			"type":        "User",
			"url":         "https://github.com/Codertocat",
		},
	}

	dest := mapDiscussion(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
