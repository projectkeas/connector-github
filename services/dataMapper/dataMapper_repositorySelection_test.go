package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestRepositorySelectionMapping(t *testing.T) {
	src := map[string]interface{}{
		"selection": "selected",
		"added": []map[string]interface{}{
			{
				"id":        186853007,
				"node_id":   "MDEwOlJlcG9zaXRvcnkxODY4NTMwMDc=",
				"name":      "Space",
				"full_name": "Codertocat/Space",
				"private":   false,
			},
		},
		"removed": []map[string]interface{}{
			{
				"id":        186853007,
				"node_id":   "MDEwOlJlcG9zaXRvcnkxODY4NTMwMDc=",
				"name":      "Space",
				"full_name": "Codertocat/Space2",
				"private":   false,
			},
		},
	}
	expected := map[string]interface{}{
		"selection": "selected",
		"added": []map[string]interface{}{
			{
				"id":        186853007,
				"name":      "Space",
				"fullName":  "Codertocat/Space",
				"isPrivate": false,
			},
		},
		"removed": []map[string]interface{}{
			{
				"id":        186853007,
				"name":      "Space",
				"fullName":  "Codertocat/Space2",
				"isPrivate": false,
			},
		},
	}

	dest := mapRepositorySelection(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
