package types

// Success
type V1AiPhotoEditorCreateResponse struct {
	// The amount of credits deducted from your account to generate the image. We charge credits right when the request is made.
	//
	// If an error occurred while generating the image(s), credits will be refunded and this field will be updated to include the refund.
	CreditsCharged int `json:"credits_charged"`
	// Deprecated: Previously represented the number of frames (original name of our credit system) used for image generation. Use 'credits_charged' instead.
	FrameCost int `json:"frame_cost"`
	// Unique ID of the image. This value can be used in the [get image project API](https://docs.magichour.ai/api-reference/image-projects/get-image-details) to fetch additional details such as status
	Id string `json:"id"`
}
