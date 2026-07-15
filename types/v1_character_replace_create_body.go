package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1CharacterReplaceCreateBody
type V1CharacterReplaceCreateBody struct {
	// Source video and reference character image for the job.
	Assets V1CharacterReplaceCreateBodyAssets `json:"assets"`
	// End time of your clip (seconds). Must be greater than start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output video resolution. Defaults to 480p, the lowest resolution available on your plan.
	Resolution nullable.Nullable[V1CharacterReplaceCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64] `json:"start_seconds,omitempty"`
	// Optional style controls for replace vs animate mode and subject selection.
	Style nullable.Nullable[V1CharacterReplaceCreateBodyStyle] `json:"style,omitempty"`
}
