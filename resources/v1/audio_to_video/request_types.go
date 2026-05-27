package audio_to_video

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the audio file and an optional reference image.
	Assets types.V1AudioToVideoCreateBodyAssets `json:"assets"`
	// End time of your clip (seconds). Must be greater than start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output video resolution. Defaults to `720p` on paid tiers and `480p` on free tiers.
	Resolution nullable.Nullable[types.V1AudioToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64] `json:"start_seconds,omitempty"`
	// Attributes used to dictate the style of the output
	Style nullable.Nullable[types.V1AudioToVideoCreateBodyStyle] `json:"style,omitempty"`
}
