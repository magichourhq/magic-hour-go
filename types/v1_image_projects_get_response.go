package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Success
type V1ImageProjectsGetResponse struct {
	CreatedAt string `json:"created_at"`
	// The amount of credits deducted from your account to generate the image. We charge credits right when the request is made.
	//
	// If an error occurred while generating the image(s), credits will be refunded and this field will be updated to include the refund.
	CreditsCharged int                                       `json:"credits_charged"`
	Downloads      []V1ImageProjectsGetResponseDownloadsItem `json:"downloads"`
	// Whether this resource is active. If false, it is deleted.
	Enabled bool `json:"enabled"`
	// In the case of an error, this object will contain the error encountered during video render
	Error nullable.Nullable[V1ImageProjectsGetResponseError] `json:"error,omitempty"`
	// Unique ID of the image. Use it with the [Get image Project API](https://docs.magichour.ai/api-reference/image-projects/get-image-details) to fetch status and downloads.
	Id string `json:"id"`
	// Number of images generated
	ImageCount int `json:"image_count"`
	// The name of the image.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The status of the image.
	Status V1ImageProjectsGetResponseStatusEnum `json:"status"`
	// Deprecated: Previously represented the number of frames (original name of our credit system) used for image generation. Use 'credits_charged' instead.
	TotalFrameCost int `json:"total_frame_cost"`
	// The type of the image project. Possible values are AI_HEADSHOT, AI_IMAGE, IMAGE_UPSCALER, FACE_SWAP, PHOTO_EDITOR, QR_CODE, BACKGROUND_REMOVER, CLOTHES_CHANGER, AI_MEME, FACE_EDITOR, PHOTO_COLORIZER, AI_GIF, AI_SELFIE, AI_IMAGE_EDITOR
	Type string `json:"type"`
}
