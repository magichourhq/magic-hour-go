
package types
import (
nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1AiImageUpscalerBody
type PostV1AiImageUpscalerBody struct {
    // Provide the assets for upscaling
    Assets PostV1AiImageUpscalerBodyAssets `json:"assets"`
    // The name of image
    Name nullable.Nullable[string] `json:"name,omitempty"`
    // How much to scale the image. Must be either 2 or 4
    ScaleFactor float64 `json:"scale_factor"`
    Style PostV1AiImageUpscalerBodyStyle `json:"style"`
}



