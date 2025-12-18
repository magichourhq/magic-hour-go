package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// The art style to use for image generation.
type V1AiImageGeneratorCreateBodyStyle struct {
	// The prompt used for the image(s).
	Prompt string `json:"prompt"`
	// Controls the quality of the generated image. Defaults to 'standard' if not specified.
	//
	// **Options:**
	// - `standard` - Standard quality generation. Cost: 5 credits per image.
	// - `pro` - Pro quality generation with enhanced details and quality. Cost: 30 credits per image.
	//
	// Note: Pro mode is available for users on Creator, Pro, or Business tier.
	QualityMode nullable.Nullable[V1AiImageGeneratorCreateBodyStyleQualityModeEnum] `json:"quality_mode,omitempty"`
	// The art style to use for image generation. Defaults to 'general' if not provided.
	Tool nullable.Nullable[V1AiImageGeneratorCreateBodyStyleToolEnum] `json:"tool,omitempty"`
}
