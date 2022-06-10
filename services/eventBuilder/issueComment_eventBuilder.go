package eventBuilder

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

func mapIssueComment(payload map[string]interface{}, id string) (cloudevents.Event, cee.ValidationError) {
	event := cloudevents.NewEvent(Version)

	event.SetID(id)
	event.SetSource(lookupValueFromPayloadPath(payload, "issue", "url"))
	event.SetType("com.github.issue_comment." + lookupValueFromPayloadPath(payload, "action"))
	event.SetSubject(lookupValueFromPayloadPath(payload, "comment", "id"))
	event.SetTime(parseTime(lookupValueFromPayloadPath(payload, "comment", "updated_at")))
	event.SetData(*cloudevents.StringOfApplicationJSON(), payload)

	err := event.Validate()
	if err != nil {
		return event, err.(cee.ValidationError)
	}

	return event, nil
}
