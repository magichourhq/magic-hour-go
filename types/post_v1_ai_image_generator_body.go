package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1AiImageGeneratorBody
type PostV1AiImageGeneratorBody struct {
	// number to images to generate
	ImageCount int `json:"image_count"`
	// The name of image
	Name        nullable.Nullable[string]                 `json:"name,omitempty"`
	Orientation PostV1AiImageGeneratorBodyOrientationEnum `json:"orientation"`
	Style       PostV1AiImageGeneratorBodyStyle           `json:"style"`
}
