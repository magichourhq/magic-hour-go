package types

// How often the subscription is billed. `null` if unknown.
type V1AccountListResponseSubscriptionBillingIntervalEnum string

const (
	V1AccountListResponseSubscriptionBillingIntervalEnumMonth V1AccountListResponseSubscriptionBillingIntervalEnum = "month"
	V1AccountListResponseSubscriptionBillingIntervalEnumYear  V1AccountListResponseSubscriptionBillingIntervalEnum = "year"
)
