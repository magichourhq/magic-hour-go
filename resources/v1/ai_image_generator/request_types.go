package ai_image_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// number to images to generate
	ImageCount int `json:"image_count"`
	// The name of image
	Name        nullable.Nullable[string]                         `json:"name,omitempty"`
	Orientation types.V1AiImageGeneratorCreateBodyOrientationEnum `json:"orientation"`
	Style       types.V1AiImageGeneratorCreateBodyStyle           `json:"style"`
}
