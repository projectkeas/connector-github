package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestLabelMapping(t *testing.T) {
	src := map[string]interface{}{
		"color":   "d73a4a",
		"default": true,
		"id":      1362934389,
		"name":    "bug",
		"node_id": "MDU6TGFiZWwxMzYyOTM0Mzg5",
		"url":     "https://api.github.com/repos/Codertocat/Hello-World/labels/bug",
	}
	expected := map[string]interface{}{
		"color":   "d73a4a",
		"default": true,
		"id":      1362934389,
		"name":    "bug",
	}

	dest := mapLabel(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
