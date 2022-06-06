package dataMapper

func mapCheckRun(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndTransformElements(src, "pull_requests", data, func(input map[string]interface{}) map[string]interface{} {
		temp := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(input, "url", temp)
		addToCollectionIfFoundAndNotNil(input, "id", temp)
		addToCollectionIfFoundAndNotNil(input, "number", temp)

		addToCollectionIfFoundAndNotNilAndTransformMap(input, "head", temp, func(headData map[string]interface{}) map[string]interface{} {
			headTemp := &map[string]interface{}{}

			addToCollectionIfFoundAndNotNil(headData, "ref", headTemp)
			addToCollectionIfFoundAndNotNil(headData, "sha", headTemp)
			addToCollectionIfFoundAndNotNil(headData, "repo", headTemp)

			return *headTemp
		})

		addToCollectionIfFoundAndNotNilAndTransformMap(input, "base", temp, func(baseData map[string]interface{}) map[string]interface{} {
			baseTemp := &map[string]interface{}{}

			addToCollectionIfFoundAndNotNil(baseData, "ref", baseTemp)
			addToCollectionIfFoundAndNotNil(baseData, "sha", baseTemp)
			addToCollectionIfFoundAndNotNil(baseData, "repo", baseTemp)

			return *baseTemp
		})

		return *temp
	})

	addToCollectionIfFoundAndNotNilAndTransformMap(src, "output", data, func(output map[string]interface{}) map[string]interface{} {
		dataOutput := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(output, "title", dataOutput)
		addToCollectionIfFoundAndNotNil(output, "summary", dataOutput)
		addToCollectionIfFoundAndNotNil(output, "text", dataOutput)
		addToCollectionIfFoundAndNotNil(output, "annotations_count", dataOutput)
		addToCollectionIfFoundAndNotNil(output, "annotations_url", dataOutput)

		return *dataOutput
	})

	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "head_sha", data)
	addToCollectionIfFoundAndNotNil(src, "external_id", data)
	addToCollectionIfFoundAndNotNil(src, "html_url", data)
	addToCollectionIfFoundAndNotNil(src, "details_url", data)
	addToCollectionIfFoundAndNotNil(src, "conclusion", data)
	addToCollectionIfFoundAndNotNil(src, "status", data)
	addToCollectionIfFoundAndNotNil(src, "started_at", data)
	addToCollectionIfFoundAndNotNil(src, "completed_at", data)
	addToCollectionIfFoundAndNotNil(src, "name", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "check_suite", data, mapCheckSuite)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "app", data, mapApp)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "deployment", data, mapDeployment)
	addToCollectionIfFoundAndNotNil(src, "created_at", data)

	return *data
}
