package dataMapper

import (
	"testing"

	cmp "github.com/google/go-cmp/cmp"
)

func TestMarketplacePurchaseMapping(t *testing.T) {
	src := map[string]interface{}{
		"account": map[string]interface{}{
			"type":                       "Organization",
			"id":                         18404719,
			"node_id":                    "MDEyOk9yZ2FuaXphdGlvbjE=",
			"login":                      "username",
			"organization_billing_email": "username@email.com",
		},
		"billing_cycle":      "monthly",
		"unit_count":         1,
		"on_free_trial":      false,
		"free_trial_ends_on": "2017-11-05T00:00:00+00:00",
		"next_billing_date":  "2017-11-05T00:00:00+00:00",
		"plan": map[string]interface{}{
			"id":                     435,
			"name":                   "Basic Plan",
			"description":            "Basic Plan",
			"monthly_price_in_cents": 1000,
			"yearly_price_in_cents":  10000,
			"price_model":            "per-unit",
			"has_free_trial":         true,
			"unit_name":              "seat",
			"bullets": []string{
				"Is Basic",
				"Because Basic",
			},
		},
	}
	expected := map[string]interface{}{
		"account": map[string]interface{}{
			"type":                     "Organization",
			"id":                       18404719,
			"login":                    "username",
			"organizationBillingEmail": "username@email.com",
		},
		"billingCycle":    "monthly",
		"unitCount":       1,
		"onFreeTrial":     false,
		"freeTrialEndsOn": "2017-11-05T00:00:00+00:00",
		"nextBillingDate": "2017-11-05T00:00:00+00:00",
		"plan": map[string]interface{}{
			"id":                  435,
			"name":                "Basic Plan",
			"description":         "Basic Plan",
			"monthlyPriceInCents": 1000,
			"yearlyPriceInCents":  10000,
			"pricingModel":        "per-unit",
			"hasFreeTrial":        true,
			"unitName":            "seat",
			"bullets": []string{
				"Is Basic",
				"Because Basic",
			},
		},
	}

	dest := mapMarketplacePurchase(src)

	if diff := cmp.Diff(expected, dest); diff != "" {
		t.Errorf("Returned value did not match expected. Difference (-want +got):\n%s", diff)
	}
}
