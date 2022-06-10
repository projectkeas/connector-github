package eventBuilder

import (
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

func mapGollum(payload map[string]interface{}, id string) (cloudevents.Event, cee.ValidationError) {
	event := cloudevents.NewEvent(Version)

	el := payload["pages"]
	// TODO :: Verify with Github
	tempPayload := el.([]interface{})[0].(map[string]interface{})

	event.SetID(id)
	event.SetSource(lookupValueFromPayloadPath(payload, "repository", "url"))
	event.SetType("com.github.gollum." + lookupValueFromPayloadPath(tempPayload, "action"))
	event.SetSubject(lookupValueFromPayloadPath(tempPayload, "page_name"))
	event.SetTime(time.Now().UTC())
	event.SetData(*cloudevents.StringOfApplicationJSON(), payload)

	err := event.Validate()
	if err != nil {
		return event, err.(cee.ValidationError)
	}

	return event, nil
}
