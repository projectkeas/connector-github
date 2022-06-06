package dataMapper

func mapApp(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "description", data)
	addToCollectionIfFoundAndNotNil(src, "events", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "name", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "owner", data, mapLogin)
	addToCollectionIfFoundAndNotNil(src, "permissions", data)
	addToCollectionIfFoundAndNotNil(src, "slug", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNil(src, "html_url", data)

	return *data
}
