package dataMapper

func mapAlert(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndTransformElements(src, "instances", data, func(input map[string]interface{}) map[string]interface{} {
		temp := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(input, "ref", temp)
		addToCollectionIfFoundAndNotNil(input, "analysis_key", temp)
		addToCollectionIfFoundAndNotNil(input, "environment", temp)
		addToCollectionIfFoundAndNotNil(input, "state", temp)

		return *temp
	})

	addToCollectionIfFoundAndNotNilAndTransformMap(src, "rule", data, func(input map[string]interface{}) map[string]interface{} {
		temp := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(input, "id", temp)
		addToCollectionIfFoundAndNotNil(input, "severity", temp)
		addToCollectionIfFoundAndNotNil(input, "description", temp)
		addToCollectionIfFoundAndNotNil(input, "full_description", temp)
		addToCollectionIfFoundAndNotNil(input, "tags", temp)
		addToCollectionIfFoundAndNotNil(input, "help", temp)

		return *temp
	})

	addToCollectionIfFoundAndNotNilAndTransformMap(src, "tool", data, func(input map[string]interface{}) map[string]interface{} {
		temp := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(input, "name", temp)
		addToCollectionIfFoundAndNotNil(input, "version", temp)

		return *temp
	})

	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "number", data)
	addToCollectionIfFoundAndNotNil(src, "html_url", data)
	addToCollectionIfFoundAndNotNil(src, "state", data)
	addToCollectionIfFoundAndNotNil(src, "fixed_at", data)
	addToCollectionIfFoundAndNotNil(src, "dismissed_by", data)
	addToCollectionIfFoundAndNotNil(src, "dismissed_at", data)
	addToCollectionIfFoundAndNotNil(src, "dismissed_reason", data)
	addToCollectionIfFoundAndNotNil(src, "ref", data)
	addToCollectionIfFoundAndNotNil(src, "commit_id", data)

	return *data
}
