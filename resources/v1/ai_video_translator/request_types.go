package ai_video_translator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Source video for the translation job.
	Assets types.V1AiVideoTranslatorCreateBodyAssets `json:"assets"`
	// End time of your clip (seconds). Must be greater than start_seconds. The clip must be 1-30 seconds long.
	EndSeconds float64 `json:"end_seconds"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output video resolution. Defaults to 480p. 720p and 1080p require a paid plan.
	Resolution nullable.Nullable[types.V1AiVideoTranslatorCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64] `json:"start_seconds,omitempty"`
	// Language to translate the video's speech into.
	TargetLanguage types.V1AiVideoTranslatorCreateBodyTargetLanguageEnum `json:"target_language"`
}
