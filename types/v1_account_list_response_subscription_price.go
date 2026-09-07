package types

// V1AccountListResponseSubscriptionPrice
type V1AccountListResponseSubscriptionPrice struct {
	// Price charged per billing interval, in the smallest unit of the currency (e.g. 4900 is $49.00 for `usd`). Discounts are not applied.
	Amount int `json:"amount"`
	// Three-letter ISO 4217 currency code, lowercase.
	Currency string `json:"currency"`
}
