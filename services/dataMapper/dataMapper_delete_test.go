package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestDeleteMapping(t *testing.T) {
	src := map[string]interface{}{
		"ref":         "simple-tag",
		"ref_type":    "tag",
		"pusher_type": "user",
	}
	expected := map[string]interface{}{
		"ref":        "simple-tag",
		"refType":    "tag",
		"pusherType": "user",
	}

	dest := mapDelete(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
