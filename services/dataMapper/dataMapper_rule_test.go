package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestRuleMapping(t *testing.T) {
	src := map[string]interface{}{
		"id":                                     21796960,
		"repository_id":                          259377789,
		"name":                                   "production",
		"created_at":                             "2021-08-19T12:16:32.000-04:00",
		"updated_at":                             "2021-08-19T12:16:32.000-04:00",
		"pull_request_reviews_enforcement_level": "off",
		"required_approving_review_count":        1,
		"dismiss_stale_reviews_on_push":          false,
		"require_code_owner_review":              false,
		"authorized_dismissal_actors_only":       false,
		"ignore_approvals_from_contributors":     false,
		"required_status_checks": []string{
			"basic-CI",
		},
		"required_status_checks_enforcement_level":     "non_admins",
		"strict_required_status_checks_policy":         false,
		"signature_requirement_enforcement_level":      "off",
		"linear_history_requirement_enforcement_level": "off",
		"admin_enforced":                         false,
		"allow_force_pushes_enforcement_level":   "off",
		"allow_deletions_enforcement_level":      "off",
		"merge_queue_enforcement_level":          "off",
		"required_deployments_enforcement_level": "off",
		"required_conversation_resolution_level": "off",
		"authorized_actors_only":                 true,
		"authorized_actor_names": []string{
			"Codertocat",
		},
	}
	expected := map[string]interface{}{
		"adminEnforced":                            false,
		"allowDeletionsEnforcementLevel":           "off",
		"allowForcePushesEnforcementLevel":         "off",
		"authorizedActorNames":                     []string{"Codertocat"},
		"authorizedActorsOnly":                     true,
		"authorizedDismissalActorsOnly":            false,
		"createdAt":                                "2021-08-19T12:16:32.000-04:00",
		"dismissStaleReviewsOnPush":                false,
		"id":                                       21796960,
		"ignoreApprovalsFromContributors":          false,
		"linearHistoryRequirementEnforcementLevel": "off",
		"mergeQueueEnforcementLevel":               "off",
		"name":                                     "production",
		"pullRequestReviewsEnforcementLevel":       "off",
		"repositoryId":                             259377789,
		"requireCodeOwnerReview":                   false,
		"requiredApprovingReviewCount":             1,
		"requiredConversationResolutionLevel":      "off",
		"requiredDeploymentsEnforcementLevel":      "off",
		"requiredStatusChecks":                     []string{"basic-CI"},
		"requiredStatusChecksEnforcementLevel":     "non_admins",
		"signatureRequirementEnforcementLevel":     "off",
		"strictRequiredStatusChecksPolicy":         false,
		"updatedAt":                                "2021-08-19T12:16:32.000-04:00",
	}

	dest := mapRule(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
