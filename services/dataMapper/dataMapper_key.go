package dataMapper

func mapDeployKey(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "key", data)
	addToCollectionIfFoundAndNotNil(src, "url", data)
	addToCollectionIfFoundAndNotNil(src, "title", data)
	addToCollectionIfFoundAndNotNil(src, "verified", data)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "read_only", data)

	return *data
}
