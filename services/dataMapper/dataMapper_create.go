package dataMapper

func mapCreate(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "ref", data)
	addToCollectionIfFoundAndNotNil(src, "ref_type", data)
	addToCollectionIfFoundAndNotNil(src, "master_branch", data)
	addToCollectionIfFoundAndNotNil(src, "description", data)
	addToCollectionIfFoundAndNotNil(src, "pusher_type", data)

	return *data
}
