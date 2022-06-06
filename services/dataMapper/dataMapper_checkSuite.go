package dataMapper

func mapCheckSuite(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "after", data)
	addToCollectionIfFoundAndNotNil(src, "before", data)
	addToCollectionIfFoundAndNotNil(src, "conclusion", data)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "head_branch", data)
	addToCollectionIfFoundAndNotNil(src, "head_sha", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "status", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "head_commit", data, mapCommit)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "app", data, mapApp)

	return *data
}
