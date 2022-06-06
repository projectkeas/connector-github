package dataMapper

func mapOrganization(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "avatar_url", data)
	addToCollectionIfFoundAndNotNil(src, "description", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "login", data)

	return *data
}
