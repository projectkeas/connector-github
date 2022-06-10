package eventBuilder

import (
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

func mapTeamAdd(payload map[string]interface{}, id string) (cloudevents.Event, cee.ValidationError) {
	event := cloudevents.NewEvent(Version)

	event.SetID(id)
	event.SetSource(lookupValueFromPayloadPath(payload, "repository", "url"))
	event.SetType("com.github.team_add." + lookupValueFromPayloadPath(payload, "action"))
	event.SetSubject(lookupValueFromPayloadPath(payload, "team", "id"))
	event.SetTime(time.Now().UTC())
	event.SetData(*cloudevents.StringOfApplicationJSON(), payload)

	err := event.Validate()
	if err != nil {
		return event, err.(cee.ValidationError)
	}

	return event, nil
}
