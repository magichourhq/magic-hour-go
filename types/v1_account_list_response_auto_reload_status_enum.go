package types

// State of Auto Reload on the account.
// - `disabled`: off or never set up.
// - `active`: on, with no refill processing or blocked.
// - `auto_reload_pending`: a refill is processing. Check again shortly.
// - `payment_failed`: the last refill payment failed and paused Auto Reload until the payment method is updated.
// - `spending_limit_reached`: the monthly limit blocked the last refill.
//
// After an `insufficient_credits` rejection, resubmit the same request once `credits` reaches that rejection's `credits_required`.
type V1AccountListResponseAutoReloadStatusEnum string

const (
	V1AccountListResponseAutoReloadStatusEnumActive               V1AccountListResponseAutoReloadStatusEnum = "active"
	V1AccountListResponseAutoReloadStatusEnumAutoReloadPending    V1AccountListResponseAutoReloadStatusEnum = "auto_reload_pending"
	V1AccountListResponseAutoReloadStatusEnumDisabled             V1AccountListResponseAutoReloadStatusEnum = "disabled"
	V1AccountListResponseAutoReloadStatusEnumPaymentFailed        V1AccountListResponseAutoReloadStatusEnum = "payment_failed"
	V1AccountListResponseAutoReloadStatusEnumSpendingLimitReached V1AccountListResponseAutoReloadStatusEnum = "spending_limit_reached"
)
