package eventBuilder

import (
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

func mapStar(payload map[string]interface{}, id string) (cloudevents.Event, cee.ValidationError) {
	event := cloudevents.NewEvent(Version)

	t := time.Now().UTC()

	sat := lookupValueFromPayloadPath(payload, "starred_at")
	if sat != "" {
		t = parseTime(sat)
	}

	event.SetID(id)
	event.SetSource(lookupValueFromPayloadPath(payload, "repository", "url"))
	event.SetType("com.github.star." + lookupValueFromPayloadPath(payload, "action"))
	event.SetTime(t)
	event.SetData(*cloudevents.StringOfApplicationJSON(), payload)

	err := event.Validate()
	if err != nil {
		return event, err.(cee.ValidationError)
	}

	return event, nil
}
