package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiImageUpscalerCreateBody
type V1AiImageUpscalerCreateBody struct {
	// Provide the assets for upscaling
	Assets V1AiImageUpscalerCreateBodyAssets `json:"assets"`
	// The name of image. This value is mainly used for your own identification of the image.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// How much to scale the image. Must be either 2 or 4.
	//
	// Note: 4x upscale is only available on Creator, Pro, or Business tier.
	ScaleFactor float64                          `json:"scale_factor"`
	Style       V1AiImageUpscalerCreateBodyStyle `json:"style"`
}
