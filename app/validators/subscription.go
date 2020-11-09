package validators

import (
	"subalogue/db"
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

	return valid, paramErrors
}
