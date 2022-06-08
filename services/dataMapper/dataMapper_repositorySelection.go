package dataMapper

func mapRepositorySelection(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "selection", data)
	addToCollectionIfFoundAndTransformElements(src, "added", data, mapRepositorySelectionInternal)
	addToCollectionIfFoundAndTransformElements(src, "removed", data, mapRepositorySelectionInternal)

	return *data
}

func mapRepositorySelectionInternal(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "full_name", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "private", data)
	addToCollectionIfFoundAndNotNil(src, "name", data)

	return *data
}
