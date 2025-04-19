package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1TextToVideoCreateBodyStyle
type V1TextToVideoCreateBodyStyle struct {
	// The prompt used for the video.
	Prompt string `json:"prompt"`
	// * `quick` - Fastest option for rapid results. Takes ~3 minutes per 5s of video.
	// *  `studio` - Polished visuals with longer runtime. Takes ~8.5 minutes per 5s of video.
	QualityMode nullable.Nullable[V1TextToVideoCreateBodyStyleQualityModeEnum] `json:"quality_mode,omitempty"`
}
