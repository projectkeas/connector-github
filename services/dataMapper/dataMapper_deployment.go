package dataMapper

func mapDeployment(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "creator", data, mapLogin)
	addToCollectionIfFoundAndNotNil(src, "description", data)
	addToCollectionIfFoundAndNotNil(src, "environment", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "original_environment", data)
	addToCollectionIfFoundAndNotNil(src, "payload", data)
	addToCollectionIfFoundAndNotNil(src, "ref", data)
	addToCollectionIfFoundAndNotNil(src, "sha", data)
	addToCollectionIfFoundAndNotNil(src, "task", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)

	return *data
}
