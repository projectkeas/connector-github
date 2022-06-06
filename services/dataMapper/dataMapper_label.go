package dataMapper

func mapLabel(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "color", data)
	addToCollectionIfFoundAndNotNil(src, "default", data)
	addToCollectionIfFoundAndNotNil(src, "description", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "name", data)

	return *data
}
