package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

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
type V1AutoSubtitleGeneratorCreateBodyStyle struct {
	// Custom subtitle configuration.
	CustomConfig nullable.Nullable[V1AutoSubtitleGeneratorCreateBodyStyleCustomConfig] `json:"custom_config,omitempty"`
	// Preset subtitle templates. Please visit https://magichour.ai/create/auto-subtitle-generator to see the style of the existing templates.
	Template nullable.Nullable[V1AutoSubtitleGeneratorCreateBodyStyleTemplateEnum] `json:"template,omitempty"`
}
