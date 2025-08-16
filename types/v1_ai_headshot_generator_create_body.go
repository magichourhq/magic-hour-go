package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiHeadshotGeneratorCreateBody
type V1AiHeadshotGeneratorCreateBody struct {
	// Provide the assets for headshot photo
	Assets V1AiHeadshotGeneratorCreateBodyAssets `json:"assets"`
	// The name of image. This value is mainly used for your own identification of the image.
	Name  nullable.Nullable[string]                               `json:"name,omitempty"`
	Style nullable.Nullable[V1AiHeadshotGeneratorCreateBodyStyle] `json:"style,omitempty"`
}
