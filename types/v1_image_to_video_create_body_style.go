package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Attributed used to dictate the style of the output
type V1ImageToVideoCreateBodyStyle struct {
	// Deprecated: Please use `resolution` instead. For backward compatibility,
	// * `false` maps to 720p resolution
	// * `true` maps to 1080p resolution
	//
	// This field will be removed in a future version. Use the `resolution` field to directly specify the resolution.
	HighQuality nullable.Nullable[bool] `json:"high_quality,omitempty"`
	// The prompt used for the video.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
	// DEPRECATED: Please use `resolution` field instead. For backward compatibility:
	// * `quick` maps to 720p resolution
	// * `studio` maps to 1080p resolution
	//
	// This field will be removed in a future version. Use the `resolution` field to directly to specify the resolution.
	QualityMode nullable.Nullable[V1ImageToVideoCreateBodyStyleQualityModeEnum] `json:"quality_mode,omitempty"`
}
