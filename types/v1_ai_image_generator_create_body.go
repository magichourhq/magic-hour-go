package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiImageGeneratorCreateBody
type V1AiImageGeneratorCreateBody struct {
	// Number of images to generate.
	ImageCount int `json:"image_count"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The orientation of the output image(s).
	Orientation V1AiImageGeneratorCreateBodyOrientationEnum `json:"orientation"`
	// The art style to use for image generation.
	Style V1AiImageGeneratorCreateBodyStyle `json:"style"`
}
