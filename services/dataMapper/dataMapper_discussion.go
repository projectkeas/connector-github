package dataMapper

func mapDiscussion(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "html_url", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "number", data)
	addToCollectionIfFoundAndNotNil(src, "title", data)
	addToCollectionIfFoundAndNotNil(src, "answer_html_url", data)
	addToCollectionIfFoundAndNotNil(src, "answer_chosen_at", data)
	addToCollectionIfFoundAndNotNil(src, "answer_chosen_by", data)
	addToCollectionIfFoundAndNotNil(src, "state", data)
	addToCollectionIfFoundAndNotNil(src, "locked", data)
	addToCollectionIfFoundAndNotNil(src, "comments", data)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNil(src, "author_association", data)
	addToCollectionIfFoundAndNotNil(src, "active_lock_reason", data)
	addToCollectionIfFoundAndNotNil(src, "body", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "user", data, mapLogin)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "category", data, func(input map[string]interface{}) map[string]interface{} {
		result := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(input, "id", result)
		addToCollectionIfFoundAndNotNil(input, "emoji", result)
		addToCollectionIfFoundAndNotNil(input, "name", result)
		addToCollectionIfFoundAndNotNil(input, "description", result)
		addToCollectionIfFoundAndNotNil(input, "created_at", result)
		addToCollectionIfFoundAndNotNil(input, "updated_at", result)
		addToCollectionIfFoundAndNotNil(input, "slug", result)
		addToCollectionIfFoundAndNotNil(input, "is_answerable", result)

		return *result
	})

	return *data
}
