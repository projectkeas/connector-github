package dataMapper

func mapRule(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "repository_id", data)
	addToCollectionIfFoundAndNotNil(src, "name", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "pull_request_reviews_enforcement_level", data)
	addToCollectionIfFoundAndNotNil(src, "required_approving_review_count", data)
	addToCollectionIfFoundAndNotNil(src, "dismiss_stale_reviews_on_push", data)
	addToCollectionIfFoundAndNotNil(src, "require_code_owner_review", data)
	addToCollectionIfFoundAndNotNil(src, "authorized_dismissal_actors_only", data)
	addToCollectionIfFoundAndNotNil(src, "ignore_approvals_from_contributors", data)
	addToCollectionIfFoundAndNotNil(src, "required_status_checks", data)
	addToCollectionIfFoundAndNotNil(src, "required_status_checks_enforcement_level", data)
	addToCollectionIfFoundAndNotNil(src, "strict_required_status_checks_policy", data)
	addToCollectionIfFoundAndNotNil(src, "signature_requirement_enforcement_level", data)
	addToCollectionIfFoundAndNotNil(src, "linear_history_requirement_enforcement_level", data)
	addToCollectionIfFoundAndNotNil(src, "admin_enforced", data)
	addToCollectionIfFoundAndNotNil(src, "allow_force_pushes_enforcement_level", data)
	addToCollectionIfFoundAndNotNil(src, "allow_deletions_enforcement_level", data)
	addToCollectionIfFoundAndNotNil(src, "merge_queue_enforcement_level", data)
	addToCollectionIfFoundAndNotNil(src, "required_deployments_enforcement_level", data)
	addToCollectionIfFoundAndNotNil(src, "required_conversation_resolution_level", data)
	addToCollectionIfFoundAndNotNil(src, "authorized_actors_only", data)
	addToCollectionIfFoundAndNotNil(src, "authorized_actor_names", data)

	return *data
}
