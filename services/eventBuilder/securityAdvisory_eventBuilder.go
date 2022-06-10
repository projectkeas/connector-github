package eventBuilder

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

func mapSecurityAdvisory(payload map[string]interface{}, id string) (cloudevents.Event, cee.ValidationError) {
	event := cloudevents.NewEvent(Version)

	event.SetID(id)
	event.SetSource("github.com")
	event.SetType("com.github.security_advisory." + lookupValueFromPayloadPath(payload, "action"))
	event.SetSubject(lookupValueFromPayloadPath(payload, "security_advisory", "ghsa_id"))
	event.SetTime(parseTime(lookupValueFromPayloadPath(payload, "security_advisory", "updated_at")))
	event.SetData(*cloudevents.StringOfApplicationJSON(), payload)

	err := event.Validate()
	if err != nil {
		return event, err.(cee.ValidationError)
	}

	return event, nil
}
