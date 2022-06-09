package dataMapper

func mapHook(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "type", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "name", data)
	addToCollectionIfFoundAndNotNil(src, "active", data)
	addToCollectionIfFoundAndNotNil(src, "events", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "config", data, func(input map[string]interface{}) map[string]interface{} {
		temp := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(input, "content_type", temp)
		addToCollectionIfFoundAndNotNil(input, "insecure_ssl", temp)
		addToCollectionIfFoundAndNotNil(input, "url", temp)

		return *temp
	})

	return *data
}
