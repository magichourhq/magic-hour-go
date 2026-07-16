package character_replace

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Source video and reference character image for the job.
	Assets types.V1CharacterReplaceCreateBodyAssets `json:"assets"`
	// End time of your clip (seconds). Must be greater than start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output video resolution. Defaults to 480p, the lowest resolution available on your plan.
	Resolution nullable.Nullable[types.V1CharacterReplaceCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64] `json:"start_seconds,omitempty"`
	// Optional style controls for replace vs animate mode and subject selection.
	Style nullable.Nullable[types.V1CharacterReplaceCreateBodyStyle] `json:"style,omitempty"`
}
