package dataMapper

func mapPages(src []map[string]interface{}) []map[string]interface{} {
	data := &[]map[string]interface{}{}

	for _, value := range src {
		temp := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(value, "page_name", temp)
		addToCollectionIfFoundAndNotNil(value, "title", temp)
		addToCollectionIfFoundAndNotNil(value, "summary", temp)
		addToCollectionIfFoundAndNotNil(value, "action", temp)
		addToCollectionIfFoundAndNotNil(value, "sha", temp)
		addToCollectionIfFoundAndNotNil(value, "html_url", temp)

		if len(*temp) > 0 {
			(*data) = append(*data, *temp)
		}
	}

	return *data
}
