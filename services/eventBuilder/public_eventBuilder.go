package eventBuilder

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

func mapPublic(payload map[string]interface{}, id string) (cloudevents.Event, cee.ValidationError) {
	event := cloudevents.NewEvent(Version)

	event.SetID(id)
	event.SetSource(lookupValueFromPayloadPath(payload, "repository", "owner", "url"))
	event.SetType("com.github.public")
	event.SetSubject(lookupValueFromPayloadPath(payload, "repository", "name"))
	event.SetTime(parseTime(lookupValueFromPayloadPath(payload, "repository", "updated_at")))
	event.SetData(*cloudevents.StringOfApplicationJSON(), payload)

	err := event.Validate()
	if err != nil {
		return event, err.(cee.ValidationError)
	}

	return event, nil
}
