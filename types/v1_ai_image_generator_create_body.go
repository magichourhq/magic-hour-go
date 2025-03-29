package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiImageGeneratorCreateBody
type V1AiImageGeneratorCreateBody struct {
	// number to images to generate
	ImageCount int `json:"image_count"`
	// The name of image
	Name        nullable.Nullable[string]                   `json:"name,omitempty"`
	Orientation V1AiImageGeneratorCreateBodyOrientationEnum `json:"orientation"`
	Style       V1AiImageGeneratorCreateBodyStyle           `json:"style"`
}
