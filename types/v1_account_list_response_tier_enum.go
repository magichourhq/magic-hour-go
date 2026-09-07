package types

// Subscription tier in effect for the account. `free` if there is no active subscription, including while a subscription is `past_due`.
type V1AccountListResponseTierEnum string

const (
	V1AccountListResponseTierEnumBusiness V1AccountListResponseTierEnum = "business"
	V1AccountListResponseTierEnumCreator  V1AccountListResponseTierEnum = "creator"
	V1AccountListResponseTierEnumFree     V1AccountListResponseTierEnum = "free"
	V1AccountListResponseTierEnumPro      V1AccountListResponseTierEnum = "pro"
)
