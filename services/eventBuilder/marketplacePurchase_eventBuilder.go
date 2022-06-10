package eventBuilder

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cee "github.com/cloudevents/sdk-go/v2/event"
)

func mapMarketplacePurchase(payload map[string]interface{}, id string) (cloudevents.Event, cee.ValidationError) {
	event := cloudevents.NewEvent(Version)

	event.SetID(id)
	event.SetSource(lookupValueFromPayloadPath(payload, "sender", "url"))
	event.SetType("com.github.marketplace_purchase." + lookupValueFromPayloadPath(payload, "action"))
	event.SetSubject(lookupValueFromPayloadPath(payload, "marketplace_purchase", "account", "login"))
	event.SetTime(parseTime(lookupValueFromPayloadPath(payload, "effective_date")))
	event.SetData(*cloudevents.StringOfApplicationJSON(), payload)

	err := event.Validate()
	if err != nil {
		return event, err.(cee.ValidationError)
	}

	return event, nil
}
