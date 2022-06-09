package dataMapper

import (
	"time"
)

type mapper func(map[string]interface{}) map[string]interface{}
type collectionMapper func([]map[string]interface{}) []map[string]interface{}

var (
	mappers map[string]mapper = map[string]mapper{
		"alert":                mapAlert,
		"app":                  mapApp,
		"blocked_user":         mapLogin,
		"changes":              remapChanges,
		"check_run":            mapCheckRun,
		"check_suite":          mapCheckSuite,
		"comment":              mapComment,
		"create":               mapCreate,
		"delete":               mapDelete,
		"deployment":           mapDeployment,
		"deployment_status":    mapDeploymentStatus,
		"discussion":           mapDiscussion,
		"enterprise":           mapEnterprise,
		"forkee":               mapRepository,
		"installation":         mapInstallation,
		"issue":                mapIssue,
		"hook":                 mapHook,
		"key":                  mapDeployKey,
		"label":                mapLabel,
		"marketplace_purchase": mapMarketplacePurchase,
		"member":               mapLogin,
		"membership":           mapMembership,
		"milestone":            mapMilestone,
		"organization":         mapOrganization,
		"pull_request":         mapPullRequest,
		"repository":           mapRepository,
		"repositorySelection":  mapRepositorySelection,
		"review":               mapReview,
		"rule":                 mapRule,
		"sender":               mapSender,
		"team":                 mapTeam,
	}
	collectionMappers = map[string]collectionMapper{
		"pages": mapPages,
	}
)

func MapWebhook(src map[string]interface{}, eventName string) map[string]interface{} {

	switch eventName {
	case "code_scanning_alert":
		prepareCodeScanningEvent(&src)
	case "create":
		prepareCreateEvent(&src)
	case "delete":
		prepareDeleteEvent(&src)
	case "push":
		preparePushEvent(&src)
	case "installation_repositories":
		prepareInstallationRepositories(&src)
	}

	dest := &map[string]interface{}{}

	for key, value := range src {
		mapper, found := mappers[key]
		if found {
			convertedValue, ok := value.(map[string]interface{})
			if ok {
				mappedValue := mapper(convertedValue)
				if len(mappedValue) > 0 {
					(*dest)[remapKey(key)] = mappedValue
				}
			}
		}

		collectionMapper, found := collectionMappers[key]
		if !found {
			continue
		}
		convertedValue, ok := value.([]map[string]interface{})
		if ok {
			mappedValue := collectionMapper(convertedValue)
			(*dest)[remapKey(key)] = mappedValue
		} else {
			convertedValue, ok := value.([]interface{})
			if ok {
				// need to proxy this back to a map[string]interface for conversion
				data := []map[string]interface{}{}
				for _, element := range convertedValue {
					el, ok := element.(map[string]interface{})
					if ok {
						data = append(data, el)
					}
				}

				if len(data) > 0 {
					mappedValue := collectionMapper(data)
					(*dest)[remapKey(key)] = mappedValue
				}
			}
		}
	}

	return *dest
}

func preparePushEvent(src *map[string]interface{}) {
	push := map[string]interface{}{}
	for key, value := range *src {
		switch key {
		case "ref":
			fallthrough
		case "before":
			fallthrough
		case "after":
			fallthrough
		case "created":
			fallthrough
		case "deleted":
			fallthrough
		case "forced":
			push[key] = value

		case "commits":
			commits := value.([]interface{})
			mappedCommits := []map[string]interface{}{}
			for _, commit := range commits {
				mappedCommits = append(mappedCommits, mapCommit(commit.(map[string]interface{})))
			}
			(*src)["commits"] = mappedCommits
		case "head_commit":
			(*src)["headCommit"] = mapCommit(value.(map[string]interface{}))
		case "pusher":
			pusher := value.(map[string]interface{})
			(*src)["pusher"] = map[string]interface{}{
				"name":  pusher["name"],
				"email": pusher["email"],
			}
		}
	}
	// store the details in a friendlier format
	(*src)["push"] = push
}

func prepareCodeScanningEvent(src *map[string]interface{}) {
	alertSrc := (*src)["alert"].(map[string]interface{})
	alertSrc["ref"] = (*src)["ref"]
	alertSrc["commit_id"] = (*src)["commit_oid"]
}

func prepareCreateEvent(src *map[string]interface{}) {
	(*src)["create"] = map[string]interface{}{
		"ref":           (*src)["ref"],
		"ref_type":      (*src)["ref_type"],
		"master_branch": (*src)["master_branch"],
		"description":   (*src)["description"],
		"pusher_type":   (*src)["pusher_type"],
	}
}

func prepareDeleteEvent(src *map[string]interface{}) {
	(*src)["delete"] = map[string]interface{}{
		"ref":         (*src)["ref"],
		"ref_type":    (*src)["ref_type"],
		"pusher_type": (*src)["pusher_type"],
	}
}

func prepareInstallationRepositories(src *map[string]interface{}) {
	(*src)["repositorySelection"] = map[string]interface{}{
		"selection": (*src)["repository_selection"],
		"added":     (*src)["repositories_added"],
		"removed":   (*src)["repositories_removed"],
	}
}

func remapKey(key string) string {
	remappedKey, found := keyRemap[key]
	if found {
		return remappedKey
	}

	return key
}

func remapChanges(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	// we only need to remap the keys as required, let the object structure stay the same
	for key, value := range src {
		(*data)[remapKey(key)] = value
	}

	return *data
}

func mapLogin(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "avatar_url", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "site_admin", data)
	addToCollectionIfFoundAndNotNil(src, "login", data)
	addToCollectionIfFoundAndNotNil(src, "slug", data)
	addToCollectionIfFoundAndNotNil(src, "type", data)
	addToCollectionIfFoundAndNotNil(src, "html_url", data)

	return *data
}

func mapRepositoryRef(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "label", data)
	addToCollectionIfFoundAndNotNil(src, "ref", data)
	addToCollectionIfFoundAndNotNil(src, "sha", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "user", data, mapLogin)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "repo", data, mapRepository)

	return *data
}

func mapCommit(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}
	addToCollectionIfFoundAndNotNil(src, "added", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "author", data, mapAuthor)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "committer", data, mapAuthor)
	addToCollectionIfFoundAndNotNil(src, "distinct", data)
	addToCollectionIfFoundAndNotNil(src, "id", data)
	addToCollectionIfFoundAndNotNil(src, "message", data)
	addToCollectionIfFoundAndNotNil(src, "modified", data)
	addToCollectionIfFoundAndNotNil(src, "removed", data)
	addToCollectionIfFoundAndNotNil(src, "timestamp", data)
	addToCollectionIfFoundAndNotNil(src, "tree_id", data)
	addToCollectionIfFoundAndNotNil(src, "url", data)

	return *data
}

func mapAuthor(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "email", data)
	addToCollectionIfFoundAndNotNil(src, "name", data)
	addToCollectionIfFoundAndNotNil(src, "username", data)

	return *data
}

func addToCollectionIfFoundAndNotNil(src map[string]interface{}, srcKey string, dest *map[string]interface{}) {
	value, found := src[srcKey]
	if found && value != nil {
		(*dest)[remapKey(srcKey)] = value
	}
}

func addToCollectionIfFoundAndNotNilAndTransformValue(src map[string]interface{}, srcKey string, dest *map[string]interface{}, transform func(interface{}) interface{}) {
	value, found := src[srcKey]
	if found && value != nil {
		(*dest)[remapKey(srcKey)] = transform(value)
	}
}

func addToCollectionIfFoundAndNotNilAndTransformMap(src map[string]interface{}, srcKey string, dest *map[string]interface{}, transform func(map[string]interface{}) map[string]interface{}) {
	value, found := src[srcKey]
	if found && value != nil {
		convertedValue, ok := value.(map[string]interface{})
		if ok {
			(*dest)[remapKey(srcKey)] = transform(convertedValue)
		}
	}
}

func addToCollectionIfFoundAndTransformElements(src map[string]interface{}, srcKey string, dest *map[string]interface{}, transform func(map[string]interface{}) map[string]interface{}) {
	value, found := src[srcKey]
	if !found || value == nil {
		return
	}
	convertedElements := []map[string]interface{}{}

	convertedValue, ok := value.([]interface{})
	if ok {
		for _, elementInf := range convertedValue {
			element, ok := elementInf.(map[string]interface{})
			if !ok {
				continue
			}

			elementData := transform(element)

			if len(elementData) > 0 {
				convertedElements = append(convertedElements, elementData)
			}
		}
	} else {
		convertedValue, ok := value.([]map[string]interface{})
		if ok {
			for _, element := range convertedValue {
				elementData := transform(element)

				if len(elementData) > 0 {
					convertedElements = append(convertedElements, elementData)
				}
			}
		}
	}

	(*dest)[remapKey(srcKey)] = convertedElements
}

func convertToDateTimeString(input interface{}) interface{} {
	convertedF, ok := input.(float64)
	if ok {
		return time.Unix(int64(convertedF), 0).UTC().Format(time.RFC3339)
	}

	convertedI, ok := input.(int64)
	if ok {
		return time.Unix(convertedI, 0).UTC().Format(time.RFC3339)
	}

	convertedI32, ok := input.(int)
	if ok {
		return time.Unix(int64(convertedI32), 0).UTC().Format(time.RFC3339)
	}

	return input
}
