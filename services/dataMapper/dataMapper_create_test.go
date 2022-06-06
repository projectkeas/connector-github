package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestCreateMapping(t *testing.T) {
	src := map[string]interface{}{
		"ref":           "simple-tag",
		"ref_type":      "tag",
		"master_branch": "master",
		"description":   "description",
		"pusher_type":   "user",
	}
	expected := map[string]interface{}{
		"ref":           "simple-tag",
		"refType":       "tag",
		"defaultBranch": "master",
		"description":   "description",
		"pusherType":    "user",
	}

	dest := mapCreate(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
