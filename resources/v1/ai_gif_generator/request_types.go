package ai_gif_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// The name of gif
	Name  nullable.Nullable[string]             `json:"name,omitempty"`
	Style types.V1AiGifGeneratorCreateBodyStyle `json:"style"`
}
