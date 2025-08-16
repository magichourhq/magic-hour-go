package ai_image_upscaler

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for upscaling
	Assets types.V1AiImageUpscalerCreateBodyAssets `json:"assets"`
	// The name of image. This value is mainly used for your own identification of the image.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// How much to scale the image. Must be either 2 or 4.
	//
	// Note: 4x upscale is only available on Creator, Pro, or Business tier.
	ScaleFactor float64                                `json:"scale_factor"`
	Style       types.V1AiImageUpscalerCreateBodyStyle `json:"style"`
}
