package ai_image_upscaler

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for upscaling
	Assets types.PostV1AiImageUpscalerBodyAssets `json:"assets"`
	// The name of image
	Name        nullable.Nullable[string]            `json:"name,omitempty"`
	ScaleFactor float64                              `json:"scale_factor"`
	Style       types.PostV1AiImageUpscalerBodyStyle `json:"style"`
}
