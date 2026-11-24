package types

// V1AccountCheckoutCreateResponse
type V1AccountCheckoutCreateResponse struct {
	// Open in a browser to complete the purchase.
	CheckoutUrl string `json:"checkout_url"`
}
