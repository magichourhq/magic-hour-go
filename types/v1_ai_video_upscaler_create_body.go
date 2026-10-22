package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiVideoUpscalerCreateBody
type V1AiVideoUpscalerCreateBody struct {
	// Source video for the upscale job.
	Assets V1AiVideoUpscalerCreateBodyAssets `json:"assets"`
	// End time of your clip (seconds). Must be greater than start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output resolution. `720p` and `1080p` set the output's short side, `2k` and `4k` its long side; the aspect ratio is kept. Requires a paid plan.
	Resolution V1AiVideoUpscalerCreateBodyResolutionEnum `json:"resolution"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64] `json:"start_seconds,omitempty"`
}
