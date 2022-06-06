package dataMapper

func mapDeploymentStatus(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "creator", data, mapLogin)
	addToCollectionIfFoundAndNotNil(src, "deployment_url", data)
	addToCollectionIfFoundAndNotNil(src, "description", data)
	addToCollectionIfFoundAndNotNil(src, "environment", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "repository_url", data)
	addToCollectionIfFoundAndNotNil(src, "state", data)
	addToCollectionIfFoundAndNotNil(src, "target_url", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)

	return *data
}
