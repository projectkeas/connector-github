package eventBuilder

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

func mapPullRequestReview(payload map[string]interface{}, id string) (cloudevents.Event, cee.ValidationError) {
	event := cloudevents.NewEvent(Version)

	event.SetID(id)
	event.SetSource(lookupValueFromPayloadPath(payload, "pull_request", "url"))
	event.SetType("com.github.pull_request_review." + lookupValueFromPayloadPath(payload, "action"))
	event.SetSubject(lookupValueFromPayloadPath(payload, "review", "id"))
	event.SetTime(parseTime(lookupValueFromPayloadPath(payload, "review", "submitted_at")))
	event.SetData(*cloudevents.StringOfApplicationJSON(), payload)

	err := event.Validate()
	if err != nil {
		return event, err.(cee.ValidationError)
	}

	return event, nil
}
