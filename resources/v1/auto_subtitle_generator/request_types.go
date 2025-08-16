package auto_subtitle_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for auto subtitle generator
	Assets types.V1AutoSubtitleGeneratorCreateBodyAssets `json:"assets"`
	// The end time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0.1, and more than the start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// The name of video. This value is mainly used for your own identification of the video.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The start time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0.
	StartSeconds float64 `json:"start_seconds"`
	// Style of the subtitle. At least one of `.style.template` or `.style.custom_config` must be provided.
	// * If only `.style.template` is provided, default values for the template will be used.
	// * If both are provided, the fields in `.style.custom_config` will be used to overwrite the fields in `.style.template`.
	// * If only `.style.custom_config` is provided, then all fields in `.style.custom_config` will be used.
	//
	// To use custom config only, the following `custom_config` params are required:
	// * `.style.custom_config.font`
	// * `.style.custom_config.text_color`
	// * `.style.custom_config.vertical_position`
	// * `.style.custom_config.horizontal_position`
	//
	Style types.V1AutoSubtitleGeneratorCreateBodyStyle `json:"style"`
}
