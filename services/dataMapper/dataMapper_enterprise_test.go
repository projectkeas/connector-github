package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestEnterpriseMapping(t *testing.T) {
	src := map[string]interface{}{
		"id":          123213,
		"slug":        "test-enterprise",
		"name":        "enterprise.com",
		"node_id":     "MDEwOkVudGVycHJpc2UxMzk1",
		"avatar_url":  "https://avatars.githubusercontent.com/b/123213?v=4",
		"description": "Test enterprise",
		"website_url": "https://enterprise.com",
		"html_url":    "https://github.com/enterprises/enterprise",
		"created_at":  "2019-11-22T12:05:26Z",
		"updated_at":  "2021-08-09T23:33:06Z",
	}
	expected := map[string]interface{}{
		"avatarUrl":   "https://avatars.githubusercontent.com/b/123213?v=4",
		"id":          123213,
		"slug":        "test-enterprise",
		"description": "Test enterprise",
		"createdAt":   "2019-11-22T12:05:26Z",
		"name":        "enterprise.com",
		"updatedAt":   "2021-08-09T23:33:06Z",
		"website":     "https://enterprise.com",
	}

	dest := mapEnterprise(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
