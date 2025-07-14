package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1TextToVideoCreateBodyStyle
type V1TextToVideoCreateBodyStyle struct {
	// The prompt used for the video.
	Prompt string `json:"prompt"`
	// DEPRECATED: Please use `resolution` field instead. For backward compatibility:
	// * `quick` maps to 720p resolution
	// * `studio` maps to 1080p resolution
	//
	// This field will be removed in a future version. Use the `resolution` field to directly to specify the resolution.
	QualityMode nullable.Nullable[V1TextToVideoCreateBodyStyleQualityModeEnum] `json:"quality_mode,omitempty"`
}
