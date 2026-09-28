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
	// Model to use. Defaults to `wan-animate`.
	//
	// * **`wan-animate`**: 480p, 720p. Supports `points` subject selection.
	// * **`kling-3.0`**: 720p, 1080p. Clips of 3–10 seconds in `replace` mode or 3–30 seconds in `animate` mode. Picks the main person automatically, so `points` are rejected.
	Model nullable.Nullable[V1CharacterReplaceCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output video resolution. Must be supported by `model`. Defaults to the lowest resolution available on your plan for that model.
	Resolution nullable.Nullable[V1CharacterReplaceCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64] `json:"start_seconds,omitempty"`
	// Optional style controls for replace vs animate mode and subject selection.
	Style nullable.Nullable[V1CharacterReplaceCreateBodyStyle] `json:"style,omitempty"`
}
