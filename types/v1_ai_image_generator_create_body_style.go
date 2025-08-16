package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// The art style to use for image generation.
type V1AiImageGeneratorCreateBodyStyle struct {
	// The prompt used for the image(s).
	Prompt string `json:"prompt"`
	// The art style to use for image generation. Defaults to 'general' if not provided.
	Tool nullable.Nullable[V1AiImageGeneratorCreateBodyStyleToolEnum] `json:"tool,omitempty"`
}
