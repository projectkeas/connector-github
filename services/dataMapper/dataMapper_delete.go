package dataMapper

func mapDelete(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "ref", data)
	addToCollectionIfFoundAndNotNil(src, "ref_type", data)
	addToCollectionIfFoundAndNotNil(src, "pusher_type", data)

	return *data
}
