package ai_image_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Number of images to generate.
	ImageCount int `json:"image_count"`
	// The name of image. This value is mainly used for your own identification of the image.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The orientation of the output image(s).
	Orientation types.V1AiImageGeneratorCreateBodyOrientationEnum `json:"orientation"`
	// The art style to use for image generation.
	Style types.V1AiImageGeneratorCreateBodyStyle `json:"style"`
}
