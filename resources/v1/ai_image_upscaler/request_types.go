package ai_image_upscaler

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for upscaling
	Assets types.V1AiImageUpscalerCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// How much to scale the image. Must be either 2 or 4.
	//
	// Note: 4x upscale is only available on Creator, Pro, or Business tier.
	ScaleFactor float64 `json:"scale_factor"`
	// Style settings for the upscale. Use `mode` (`"preserve"`, `"balanced"`, or `"creative"`). Defaults to `"balanced"`.
	Style nullable.Nullable[types.V1AiImageUpscalerCreateBodyStyle] `json:"style,omitempty"`
}
