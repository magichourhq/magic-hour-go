package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Details of the account's subscription plan. `null` if the account has no subscription, e.g. a free account, an account that only purchased credit packs, or an account on usage-based API pricing.
//
// Reflects the plan currently configured on the subscription. If a plan change is scheduled, `tier` stays on the current plan until the next payment succeeds, so `tier` and `name` can briefly disagree.
type V1AccountListResponseSubscription struct {
	// How often the subscription is billed. `null` if unknown.
	BillingInterval nullable.Nullable[V1AccountListResponseSubscriptionBillingIntervalEnum] `json:"billing_interval,omitempty"`
	// Whether the subscription is scheduled to end at `current_period_end` instead of renewing. The subscription stays usable until then.
	CancelAtPeriodEnd bool `json:"cancel_at_period_end"`
	// End of the current billing period, in ISO 8601 format. The subscription renews at this time, or ends if `cancel_at_period_end` is `true`.
	CurrentPeriodEnd nullable.Nullable[string] `json:"current_period_end,omitempty"`
	// Discount applied to the subscription. `null` if no discount is applied.
	Discount nullable.Nullable[V1AccountListResponseSubscriptionDiscount] `json:"discount,omitempty"`
	// Name of the current subscription plan, e.g. `Creator`, `Pro`, `Pro Plus`, `Business`. `null` if the plan cannot be determined. Use `tier` for a machine-readable value.
	Name  nullable.Nullable[string]              `json:"name,omitempty"`
	Price V1AccountListResponseSubscriptionPrice `json:"price"`
	// Status of the subscription.
	// - `active`: payments are up to date.
	// - `past_due`: the latest payment failed. `tier` is `free` until payment succeeds. The subscription is canceled if payment keeps failing.
	Status V1AccountListResponseSubscriptionStatusEnum `json:"status"`
}
