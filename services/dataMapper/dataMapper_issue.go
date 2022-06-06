package dataMapper

func mapIssue(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "active_lock_reason", data)
	addToCollectionIfFoundAndNotNil(src, "author_association", data)
	addToCollectionIfFoundAndNotNil(src, "body", data)
	addToCollectionIfFoundAndNotNil(src, "closed_at", data)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "labels", data)
	addToCollectionIfFoundAndNotNil(src, "locked", data)
	addToCollectionIfFoundAndNotNil(src, "milestone", data)
	addToCollectionIfFoundAndNotNil(src, "number", data)
	addToCollectionIfFoundAndNotNil(src, "state", data)
	addToCollectionIfFoundAndNotNil(src, "title", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNil(src, "html_url", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "user", data, mapLogin)
	addToCollectionIfFoundAndTransformElements(src, "labels", data, mapLabel)
	addToCollectionIfFoundAndTransformElements(src, "assignees", data, mapLogin)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "milestone", data, mapMilestone)

	return *data
}

func mapMilestone(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "html_url", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "number", data)
	addToCollectionIfFoundAndNotNil(src, "title", data)
	addToCollectionIfFoundAndNotNil(src, "description", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "creator", data, mapLogin)
	addToCollectionIfFoundAndNotNil(src, "open_issues", data)
	addToCollectionIfFoundAndNotNil(src, "closed_issues", data)
	addToCollectionIfFoundAndNotNil(src, "state", data)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNil(src, "due_on", data)
	addToCollectionIfFoundAndNotNil(src, "closed_at", data)

	return *data
}
