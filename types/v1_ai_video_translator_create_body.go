package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiVideoTranslatorCreateBody
type V1AiVideoTranslatorCreateBody struct {
	// Source video for the translation job.
	Assets V1AiVideoTranslatorCreateBodyAssets `json:"assets"`
	// End time of your clip (seconds). Must be greater than start_seconds. The clip must be 1-30 seconds long.
	EndSeconds float64 `json:"end_seconds"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Output video resolution. Defaults to 480p. 720p and 1080p require a paid plan.
	Resolution nullable.Nullable[V1AiVideoTranslatorCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds nullable.Nullable[float64] `json:"start_seconds,omitempty"`
	// Language to translate the video's speech into.
	TargetLanguage V1AiVideoTranslatorCreateBodyTargetLanguageEnum `json:"target_language"`
}
