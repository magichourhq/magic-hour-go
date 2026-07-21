package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiVideoEditorCreateBody
type V1AiVideoEditorCreateBody struct {
	// Provide the assets for video editing.
	Assets V1AiVideoEditorCreateBodyAssets `json:"assets"`
	// End time of your clip in seconds. Must be greater than `start_seconds`. Duration must be between 3 and 10 seconds.
	EndSeconds float64 `json:"end_seconds"`
	// Editing model. Defaults to `ltx-2.3` for free tier and `gemini-omni` for paid. Use `ltx-2.3` for LTX video edit.
	Model nullable.Nullable[V1AiVideoEditorCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output resolution. Defaults to `480p` for free tier and `720p` for paid. Google Omni supports 720p only; LTX-2.3 supports 480p, 720p, and 1080p.
	Resolution nullable.Nullable[V1AiVideoEditorCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64]     `json:"start_seconds,omitempty"`
	Style        V1AiVideoEditorCreateBodyStyle `json:"style"`
}
