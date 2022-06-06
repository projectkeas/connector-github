package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestAlertMapping(t *testing.T) {
	src := map[string]interface{}{
		"number":     10,
		"created_at": "2020-07-22T14:06:31Z",
		"updated_at": "2020-07-22T14:06:31Z",
		"url":        "https://api.github.com/repos/Codertocat/Hello-World/code-scanning/alerts/10",
		"html_url":   "https://github.com/Codertocat/Hello-World/security/code-scanning/10",
		"instances": []map[string]interface{}{
			{
				"ref":          "refs/heads/main",
				"analysis_key": ".github/workflows/workflow.yml:upload",
				"environment":  "{}",
				"state":        "open",
			},
		},
		"state":            "open",
		"fixed_at":         "2020-07-22T14:06:31Z",
		"dismissed_by":     "test",
		"dismissed_at":     "2020-07-22T14:06:31Z",
		"dismissed_reason": "test",
		"rule": map[string]interface{}{
			"id":               "Style/FrozenStringLiteralComment",
			"severity":         "note",
			"description":      "Add the frozen_string_literal comment to the top of files to help transition to frozen string literals by default.",
			"full_description": "Add the frozen_string_literal comment to the top of files to help transition to frozen string literals by default.",
			"tags": []string{
				"style",
			},
			"help": "Enabled by default | Safe | Supports autocorrection | VersionAdded | VersionChanged\n--- | --- | --- | --- | ---\nEnabled | No | Yes (Unsafe) | 0.36 | 0.79\n\nThis cop is designed to help you transition from mutable string literals\nto frozen string literals.\nIt will add the comment `# frozen_string_literal: true` to the top of\nfiles to enable frozen string literals. Frozen string literals may be\ndefault in future Ruby. The comment will be added below a shebang and\nencoding comment. The frozen string literal comment is only valid in\nRuby 2.3+.\n\nNote that the cop will ignore files where the comment exists but is set\nto `false` instead of `true`.\n\n### Examples\n\n#### EnforcedStyle: always (default)\n\n```ruby\n# The `always` style will always add the frozen string literal comment\n# to a file, regardless of the Ruby version or if `freeze` or `<<` are\n# called on a string literal.\n# bad\nmodule Bar\n  # ...\nend\n\n# good\n# frozen_string_literal: true\n\nmodule Bar\n  # ...\nend\n\n# good\n# frozen_string_literal: false\n\nmodule Bar\n  # ...\nend\n```\n#### EnforcedStyle: never\n\n```ruby\n# The `never` will enforce that the frozen string literal comment does\n# not exist in a file.\n# bad\n# frozen_string_literal: true\n\nmodule Baz\n  # ...\nend\n\n# good\nmodule Baz\n  # ...\nend\n```\n#### EnforcedStyle: always_true\n\n```ruby\n# The `always_true` style enforces that the frozen string literal\n# comment is set to `true`. This is a stricter option than `always`\n# and forces projects to use frozen string literals.\n# bad\n# frozen_string_literal: false\n\nmodule Baz\n  # ...\nend\n\n# bad\nmodule Baz\n  # ...\nend\n\n# good\n# frozen_string_literal: true\n\nmodule Bar\n  # ...\nend\n```\n\n### Configurable attributes\n\nName | Default value | Configurable values\n--- | --- | ---\nEnforcedStyle | `always` | `always`, `always_true`, `never`\n\n",
		},
		"tool": map[string]interface{}{
			"name": "Rubocop",
		},
	}
	expected := map[string]interface{}{
		"createdAt":       "2020-07-22T14:06:31Z",
		"dismissedAt":     "2020-07-22T14:06:31Z",
		"dismissedBy":     "test",
		"dismissedReason": "test",
		"fixedAt":         "2020-07-22T14:06:31Z",
		"number":          10,
		"rule": map[string]interface{}{
			"description":     "Add the frozen_string_literal comment to the top of files to help transition to frozen string literals by default.",
			"fullDescription": "Add the frozen_string_literal comment to the top of files to help transition to frozen string literals by default.",
			"help":            "Enabled by default | Safe | Supports autocorrection | VersionAdded | VersionChanged\n--- | --- | --- | --- | ---\nEnabled | No | Yes (Unsafe) | 0.36 | 0.79\n\nThis cop is designed to help you transition from mutable string literals\nto frozen string literals.\nIt will add the comment `# frozen_string_literal: true` to the top of\nfiles to enable frozen string literals. Frozen string literals may be\ndefault in future Ruby. The comment will be added below a shebang and\nencoding comment. The frozen string literal comment is only valid in\nRuby 2.3+.\n\nNote that the cop will ignore files where the comment exists but is set\nto `false` instead of `true`.\n\n### Examples\n\n#### EnforcedStyle: always (default)\n\n```ruby\n# The `always` style will always add the frozen string literal comment\n# to a file, regardless of the Ruby version or if `freeze` or `<<` are\n# called on a string literal.\n# bad\nmodule Bar\n  # ...\nend\n\n# good\n# frozen_string_literal: true\n\nmodule Bar\n  # ...\nend\n\n# good\n# frozen_string_literal: false\n\nmodule Bar\n  # ...\nend\n```\n#### EnforcedStyle: never\n\n```ruby\n# The `never` will enforce that the frozen string literal comment does\n# not exist in a file.\n# bad\n# frozen_string_literal: true\n\nmodule Baz\n  # ...\nend\n\n# good\nmodule Baz\n  # ...\nend\n```\n#### EnforcedStyle: always_true\n\n```ruby\n# The `always_true` style enforces that the frozen string literal\n# comment is set to `true`. This is a stricter option than `always`\n# and forces projects to use frozen string literals.\n# bad\n# frozen_string_literal: false\n\nmodule Baz\n  # ...\nend\n\n# bad\nmodule Baz\n  # ...\nend\n\n# good\n# frozen_string_literal: true\n\nmodule Bar\n  # ...\nend\n```\n\n### Configurable attributes\n\nName | Default value | Configurable values\n--- | --- | ---\nEnforcedStyle | `always` | `always`, `always_true`, `never`\n\n",
			"id":              "Style/FrozenStringLiteralComment",
			"severity":        "note",
			"tags":            []string{"style"},
		},
		"state":     "open",
		"tool":      map[string]any{"name": "Rubocop"},
		"updatedAt": "2020-07-22T14:06:31Z",
		"url":       "https://github.com/Codertocat/Hello-World/security/code-scanning/10",
	}

	dest := mapAlert(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
