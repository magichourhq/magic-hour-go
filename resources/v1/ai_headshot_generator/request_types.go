package ai_headshot_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for headshot photo
	Assets types.V1AiHeadshotGeneratorCreateBodyAssets `json:"assets"`
	// The name of image
	Name  nullable.Nullable[string]                                     `json:"name,omitempty"`
	Style nullable.Nullable[types.V1AiHeadshotGeneratorCreateBodyStyle] `json:"style,omitempty"`
}
