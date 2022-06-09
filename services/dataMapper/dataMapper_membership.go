package dataMapper

func mapMembership(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "state", data)
	addToCollectionIfFoundAndNotNil(src, "role", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "user", data, mapLogin)

	return *data
}
