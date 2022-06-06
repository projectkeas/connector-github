package dataMapper

func mapReview(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "author_association", data)
	addToCollectionIfFoundAndNotNil(src, "body", data)
	addToCollectionIfFoundAndNotNil(src, "commit_id", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "pull_request_url", data)
	addToCollectionIfFoundAndNotNil(src, "state", data)
	addToCollectionIfFoundAndNotNil(src, "submitted_at", data)
	addToCollectionIfFoundAndNotNil(src, "html_url", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "user", data, mapLogin)

	return *data
}
