package types

// Status of the subscription.
// - `active`: payments are up to date.
// - `past_due`: the latest payment failed. `tier` is `free` until payment succeeds. The subscription is canceled if payment keeps failing.
type V1AccountListResponseSubscriptionStatusEnum string

const (
	V1AccountListResponseSubscriptionStatusEnumActive  V1AccountListResponseSubscriptionStatusEnum = "active"
	V1AccountListResponseSubscriptionStatusEnumPastDue V1AccountListResponseSubscriptionStatusEnum = "past_due"
)
