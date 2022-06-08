package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestPagesMapping(t *testing.T) {
	src := []map[string]interface{}{
		{
			"page_name": "Home",
			"title":     "Home",
			"summary":   "summary",
			"action":    "edited",
			"sha":       "6bf911d3801dd1ef957fc6ade5a8d96429e7fa39",
			"html_url":  "https://github.com/Codertocat/Hello-World/wiki/Home",
		},
	}
	expected := []map[string]interface{}{
		{
			"pageName": "Home",
			"title":    "Home",
			"summary":  "summary",
			"action":   "edited",
			"sha":      "6bf911d3801dd1ef957fc6ade5a8d96429e7fa39",
			"url":      "https://github.com/Codertocat/Hello-World/wiki/Home",
		},
	}

	dest := mapPages(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
