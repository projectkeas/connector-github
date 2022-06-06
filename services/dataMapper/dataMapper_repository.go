package dataMapper

func mapRepository(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	uris := map[string]string{}
	httpsUri, ok := src["clone_url"]
	if ok {
		uris["https"] = httpsUri.(string)
	}
	sshUri, ok := src["ssh_url"]
	if ok {
		uris["ssh"] = sshUri.(string)
	}

	if len(uris) > 0 {
		(*data)["cloneUris"] = uris
	}

	addToCollectionIfFoundAndNotNil(src, "created_at", data)
	addToCollectionIfFoundAndNotNil(src, "pushed_at", data)
	addToCollectionIfFoundAndNotNil(src, "default_branch", data)
	addToCollectionIfFoundAndNotNil(src, "description", data)
	addToCollectionIfFoundAndNotNil(src, "full_name", data)
	addToCollectionIfFoundAndNotNil(src, "has_downloads", data)
	addToCollectionIfFoundAndNotNil(src, "has_issues", data)
	addToCollectionIfFoundAndNotNil(src, "has_pages", data)
	addToCollectionIfFoundAndNotNil(src, "has_projects", data)
	addToCollectionIfFoundAndNotNil(src, "has_wiki", data)
	addToCollectionIfFoundAndNotNil(src, "homepage", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "archived", data)
	addToCollectionIfFoundAndNotNil(src, "disabled", data)
	addToCollectionIfFoundAndNotNil(src, "fork", data)
	addToCollectionIfFoundAndNotNil(src, "name", data)
	addToCollectionIfFoundAndNotNil(src, "updated_at", data)
	addToCollectionIfFoundAndNotNil(src, "language", data)
	addToCollectionIfFoundAndNotNil(src, "private", data)
	addToCollectionIfFoundAndNotNil(src, "html_url", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "owner", data, mapLogin)

	return *data
}
