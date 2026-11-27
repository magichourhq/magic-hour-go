package types

// V1AccountListResponseAutoReload
type V1AccountListResponseAutoReload struct {
	// State of Auto Reload on the account.
	// - `disabled`: off or never set up.
	// - `active`: on, with no refill processing or blocked.
	// - `auto_reload_pending`: a refill is processing. Check again shortly.
	// - `payment_failed`: the last refill payment failed and paused Auto Reload until the payment method is updated.
	// - `spending_limit_reached`: the monthly limit blocked the last refill.
	//
	// After an `insufficient_credits` rejection, resubmit the same request once `credits` reaches that rejection's `credits_required`.
	Status V1AccountListResponseAutoReloadStatusEnum `json:"status"`
}
