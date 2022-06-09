package dataMapper

func mapMarketplacePurchase(src map[string]interface{}) map[string]interface{} {
	data := &map[string]interface{}{}

	addToCollectionIfFoundAndNotNil(src, "billing_cycle", data)
	addToCollectionIfFoundAndNotNil(src, "unit_count", data)
	addToCollectionIfFoundAndNotNil(src, "on_free_trial", data)
	addToCollectionIfFoundAndNotNil(src, "free_trial_ends_on", data)
	addToCollectionIfFoundAndNotNil(src, "next_billing_date", data)
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "account", data, func(input map[string]interface{}) map[string]interface{} {
		temp := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(input, "type", temp)
		addToCollectionIfFoundAndNotNil(input, "id", temp)
		addToCollectionIfFoundAndNotNil(input, "login", temp)
		addToCollectionIfFoundAndNotNil(input, "organization_billing_email", temp)

		return *temp
	})
	addToCollectionIfFoundAndNotNilAndTransformMap(src, "plan", data, func(input map[string]interface{}) map[string]interface{} {
		temp := &map[string]interface{}{}

		addToCollectionIfFoundAndNotNil(input, "id", temp)
		addToCollectionIfFoundAndNotNil(input, "name", temp)
		addToCollectionIfFoundAndNotNil(input, "description", temp)
		addToCollectionIfFoundAndNotNil(input, "monthly_price_in_cents", temp)
		addToCollectionIfFoundAndNotNil(input, "yearly_price_in_cents", temp)
		addToCollectionIfFoundAndNotNil(input, "price_model", temp)
		addToCollectionIfFoundAndNotNil(input, "has_free_trial", temp)
		addToCollectionIfFoundAndNotNil(input, "unit_name", temp)
		addToCollectionIfFoundAndNotNil(input, "bullets", temp)

		return *temp
	})

	return *data
}
