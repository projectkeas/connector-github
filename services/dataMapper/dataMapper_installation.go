package dataMapper

func mapInstallation(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "account", data, mapLogin)
	addToCollectionIfFoundAndNotNil(src, "repository_selection", data)
	addToCollectionIfFoundAndNotNil(src, "app_id", data)
	addToCollectionIfFoundAndNotNil(src, "target_id", data)
	addToCollectionIfFoundAndNotNil(src, "target_type", data)
	addToCollectionIfFoundAndNotNil(src, "html_url", data)
	addToCollectionIfFoundAndNotNil(src, "permissions", data)
	addToCollectionIfFoundAndNotNil(src, "events", data)
	addToCollectionIfFoundAndNotNilAndTransformValue(src, "created_at", data, convertToDateTimeString)
	addToCollectionIfFoundAndNotNilAndTransformValue(src, "updated_at", data, convertToDateTimeString)
	addToCollectionIfFoundAndNotNil(src, "single_file_name", data)

	return *data
}
