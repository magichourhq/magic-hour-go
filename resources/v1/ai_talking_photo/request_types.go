package ai_talking_photo

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for creating a talking photo
	Assets types.V1AiTalkingPhotoCreateBodyAssets `json:"assets"`
	// The end time of the input audio in seconds. The maximum duration allowed is 30 seconds.
	EndSeconds float64 `json:"end_seconds"`
	// The name of image
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The start time of the input audio in seconds. The maximum duration allowed is 30 seconds.
	StartSeconds float64 `json:"start_seconds"`
	// Attributes used to dictate the style of the output
	Style nullable.Nullable[types.V1AiTalkingPhotoCreateBodyStyle] `json:"style,omitempty"`
}
