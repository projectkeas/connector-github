package dataMapper

func mapComment(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "author_association", data)
	addToCollectionIfFoundAndNotNil(src, "body", data)
	addToCollectionIfFoundAndNotNil(src, "commit_id", data)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "discussion_id", data)
	addToCollectionIfFoundAndNotNil(src, "diff_hunk", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "line", data)
	addToCollectionIfFoundAndNotNil(src, "original_commit_id", data)
	addToCollectionIfFoundAndNotNil(src, "original_line", data)
	addToCollectionIfFoundAndNotNil(src, "original_position", data)
	addToCollectionIfFoundAndNotNil(src, "original_start_line", data)
	addToCollectionIfFoundAndNotNil(src, "path", data)
	addToCollectionIfFoundAndNotNil(src, "position", data)
	addToCollectionIfFoundAndNotNil(src, "pull_request_review_id", data)
	addToCollectionIfFoundAndNotNil(src, "pull_request_url", data)
	addToCollectionIfFoundAndNotNil(src, "start_line", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNil(src, "html_url", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "user", data, mapLogin)

	return *data
}
