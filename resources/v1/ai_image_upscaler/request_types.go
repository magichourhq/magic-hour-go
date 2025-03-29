package ai_image_upscaler

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for upscaling
	Assets types.V1AiImageUpscalerCreateBodyAssets `json:"assets"`
	// The name of image
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// How much to scale the image. Must be either 2 or 4
	ScaleFactor float64                                `json:"scale_factor"`
	Style       types.V1AiImageUpscalerCreateBodyStyle `json:"style"`
}
