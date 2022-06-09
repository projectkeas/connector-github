package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestMapping(t *testing.T) {
	src := map[string]interface{}{
		"type":   "Repository",
		"id":     101047067,
		"name":   "web",
		"active": true,
		"events": []string{
			"meta",
		},
		"config": map[string]interface{}{
			"content_type": "json",
			"insecure_ssl": "0",
			"url":          "http://example.com/hook",
		},
		"updated_at": "2019-04-10T03:57:12Z",
		"created_at": "2019-04-10T03:57:12Z",
	}
	expected := map[string]interface{}{
		"type":   "Repository",
		"id":     101047067,
		"name":   "web",
		"active": true,
		"events": []string{
			"meta",
		},
		"config": map[string]interface{}{
			"contentType": "json",
			"insecureSSL": "0",
			"url":         "http://example.com/hook",
		},
		"updatedAt": "2019-04-10T03:57:12Z",
		"createdAt": "2019-04-10T03:57:12Z",
	}

	dest := mapHook(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
