package validators

import (
	"github.com/shiftingphotons/subalogue/db"
)

func ValidateSubscription(subscriptionParams db.Subscription) (bool, map[string]string) {
	var paramErrors = map[string]string{}
	var valid = true

	if subscriptionParams.Name == "" {
		paramErrors["name"] = "Name should not be empty."
		valid = false
	}

	if subscriptionParams.Url == "" {
		paramErrors["url"] = "Url should not be empty."
		valid = false
	}

	// If year is 1, than BillingDate has its default value,
	// hence it wasn't passed to the endpoint (or the string was invalid)
	if subscriptionParams.BillingDate.Year() == 1 {
		paramErrors["billing_date"] = "Invalid or missing value for billing date. Use YYYY-MM-DDT00:00:00Z format."
		valid = false
	}

	if subscriptionParams.Recurring != "monthly" && subscriptionParams.Recurring != "yearly" {
		paramErrors["recurring"] = "Recurring should be either monthly or yearly"
		valid = false
	}

	return valid, paramErrors
}
