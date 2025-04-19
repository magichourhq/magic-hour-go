package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Attributed used to dictate the style of the output
type V1ImageToVideoCreateBodyStyle struct {
	// Deprecated: Please use `quality_mode` instead. For backward compatibility, setting `high_quality: true` and `quality_mode: quick` will map to `quality_mode: studio`. Note: `quality_mode: studio` offers the same quality as `high_quality: true`.
	HighQuality nullable.Nullable[bool] `json:"high_quality,omitempty"`
	// The prompt used for the video.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
	// * `quick` - Fastest option for rapid results. Takes ~3 minutes per 5s of video.
	// *  `studio` - Polished visuals with longer runtime. Takes ~8.5 minutes per 5s of video.
	QualityMode nullable.Nullable[V1ImageToVideoCreateBodyStyleQualityModeEnum] `json:"quality_mode,omitempty"`
}
