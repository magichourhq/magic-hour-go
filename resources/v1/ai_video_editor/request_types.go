package ai_video_editor

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for video editing.
	Assets types.V1AiVideoEditorCreateBodyAssets `json:"assets"`
	// End time of your clip in seconds. Must be greater than `start_seconds`. Duration must be between 3 and 10 seconds.
	EndSeconds float64 `json:"end_seconds"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64]           `json:"start_seconds,omitempty"`
	Style        types.V1AiVideoEditorCreateBodyStyle `json:"style"`
}
