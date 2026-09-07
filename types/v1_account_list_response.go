package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AccountListResponse
type V1AccountListResponse struct {
	// Credits currently available to spend. Includes subscription credits and any purchased credit packs.
	Credits int `json:"credits"`
	// Email address of the account.
	Email nullable.Nullable[string] `json:"email,omitempty"`
	// Unique ID of the account that owns the API key.
	Id string `json:"id"`
	// Details of the account's subscription plan. `null` if the account has no subscription, e.g. a free account, an account that only purchased credit packs, or an account on usage-based API pricing.
	//
	// Reflects the plan currently configured on the subscription. If a plan change is scheduled, `tier` stays on the current plan until the next payment succeeds, so `tier` and `name` can briefly disagree.
	Subscription nullable.Nullable[V1AccountListResponseSubscription] `json:"subscription,omitempty"`
	// Subscription tier in effect for the account. `free` if there is no active subscription, including while a subscription is `past_due`.
	Tier V1AccountListResponseTierEnum `json:"tier"`
}
