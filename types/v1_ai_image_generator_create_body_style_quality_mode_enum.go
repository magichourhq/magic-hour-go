package types

// DEPRECATED: Use `model` field instead for explicit model selection.
//
// Legacy quality mode mapping:
// - `standard` → `z-image-turbo` model
// - `pro` → `seedream-v4` model
//
// If model is specified, it will take precedence over the legacy quality_mode field.
type V1AiImageGeneratorCreateBodyStyleQualityModeEnum string

const (
	V1AiImageGeneratorCreateBodyStyleQualityModeEnumPro      V1AiImageGeneratorCreateBodyStyleQualityModeEnum = "pro"
	V1AiImageGeneratorCreateBodyStyleQualityModeEnumStandard V1AiImageGeneratorCreateBodyStyleQualityModeEnum = "standard"
)
