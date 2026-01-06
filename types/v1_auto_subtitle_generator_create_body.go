package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AutoSubtitleGeneratorCreateBody
type V1AutoSubtitleGeneratorCreateBody struct {
	// Provide the assets for auto subtitle generator
	Assets V1AutoSubtitleGeneratorCreateBodyAssets `json:"assets"`
	// End time of your clip (seconds). Must be greater than start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
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
	Style V1AutoSubtitleGeneratorCreateBodyStyle `json:"style"`
}
