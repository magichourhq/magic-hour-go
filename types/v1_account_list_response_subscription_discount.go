package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Discount applied to the subscription. `null` if no discount is applied.
type V1AccountListResponseSubscriptionDiscount struct {
	// Fixed amount taken off `price.amount` each billing interval, in the smallest unit of the currency. `null` if the discount is a percentage.
	AmountOff nullable.Nullable[int] `json:"amount_off,omitempty"`
	// Percentage taken off `price.amount` each billing interval. `null` if the discount is a fixed amount.
	PercentOff nullable.Nullable[float64] `json:"percent_off,omitempty"`
}
