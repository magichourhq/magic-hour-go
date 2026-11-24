package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AccountCheckoutCreateBody
type V1AccountCheckoutCreateBody struct {
	// Credits a blocked request is missing, e.g. `credit_shortfall` from a 402 response. The checkout starts with enough packs to cover it. Defaults to one pack.
	CreditShortfall nullable.Nullable[float64] `json:"credit_shortfall,omitempty"`
	// Number of credit packs the checkout starts with. Takes priority over `credit_shortfall`.
	PackCount nullable.Nullable[int] `json:"pack_count,omitempty"`
	// What to buy. `credit_pack` is a one-time credit purchase.
	Type V1AccountCheckoutCreateBodyTypeEnum `json:"type"`
}
