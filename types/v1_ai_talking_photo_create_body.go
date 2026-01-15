package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for creating a talking photo
type V1AiTalkingPhotoCreateBody struct {
	// Provide the assets for creating a talking photo
	Assets V1AiTalkingPhotoCreateBodyAssets `json:"assets"`
	// The end time of the input audio in seconds. The maximum duration allowed is 60 seconds.
	EndSeconds float64 `json:"end_seconds"`
	// Constrains the larger dimension (height or width) of the output video. Allows you to set a lower resolution than your plan's maximum if desired. The value is capped by your plan's max resolution.
	MaxResolution nullable.Nullable[int] `json:"max_resolution,omitempty"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The start time of the input audio in seconds. The maximum duration allowed is 60 seconds.
	StartSeconds float64 `json:"start_seconds"`
	// Attributes used to dictate the style of the output
	Style nullable.Nullable[V1AiTalkingPhotoCreateBodyStyle] `json:"style,omitempty"`
}
