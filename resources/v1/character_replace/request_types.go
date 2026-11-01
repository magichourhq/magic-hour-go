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
	// Model to use. Defaults to `wan-animate`.
	//
	// * **`wan-animate`**: 480p, 720p. Supports `points` subject selection.
	// * **`kling-3.0`**: 720p, 1080p. Clips of 3–10 seconds in `replace` mode or 3–30 seconds in `animate` mode. Picks the main person automatically, so `points` are rejected.
	// * **`minimax-h3`**: 480p, 720p, 1080p. Clips of 1–60 seconds in `replace` mode or 1–60 seconds in `animate` mode. Subject `points` are `replace`-only; `animate` with `selection_mode` `point` returns 422, and with `auto` or no `selection_mode` the points are ignored.
	Model nullable.Nullable[types.V1CharacterReplaceCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output video resolution. Must be supported by `model`. Defaults to the lowest resolution available on your plan for that model.
	Resolution nullable.Nullable[types.V1CharacterReplaceCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64] `json:"start_seconds,omitempty"`
	// Optional style controls for replace vs animate mode and subject selection.
	Style nullable.Nullable[types.V1CharacterReplaceCreateBodyStyle] `json:"style,omitempty"`
}
