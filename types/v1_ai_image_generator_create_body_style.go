package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// The art style to use for image generation.
type V1AiImageGeneratorCreateBodyStyle struct {
	// The prompt used for the image(s).
	Prompt string `json:"prompt"`
	// DEPRECATED: Use `model` field instead for explicit model selection.
	//
	// Legacy quality mode mapping:
	// - `standard` → `z-image-turbo` model
	// - `pro` → `seedream` model
	//
	// If model is specified, it will take precedence over the legacy quality_mode field.
	QualityMode nullable.Nullable[V1AiImageGeneratorCreateBodyStyleQualityModeEnum] `json:"quality_mode,omitempty"`
	// The art style to use for image generation. Defaults to 'general' if not provided.
	Tool nullable.Nullable[V1AiImageGeneratorCreateBodyStyleToolEnum] `json:"tool,omitempty"`
}
