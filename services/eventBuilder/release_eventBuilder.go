package eventBuilder

import (
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

func mapRelease(payload map[string]interface{}, id string) (cloudevents.Event, cee.ValidationError) {
	event := cloudevents.NewEvent(Version)

	action := lookupValueFromPayloadPath(payload, "action")
	t := time.Now().UTC()

	if action == "published" || action == "released" {
		t = parseTime(coalesce(lookupValueFromPayloadPath(payload, "release", "published_at"), lookupValueFromPayloadPath(payload, "release", "created_at")))
	} else if action == "created" || action == "prereleased" {
		t = parseTime(lookupValueFromPayloadPath(payload, "release", "created_at"))
	}

	event.SetID(id)
	event.SetSource(lookupValueFromPayloadPath(payload, "repository", "url"))
	event.SetType("com.github.release." + action)
	event.SetSubject(lookupValueFromPayloadPath(payload, "release", "id"))
	event.SetTime(t)
	event.SetData(*cloudevents.StringOfApplicationJSON(), payload)

	err := event.Validate()
	if err != nil {
		return event, err.(cee.ValidationError)
	}

	return event, nil
}
