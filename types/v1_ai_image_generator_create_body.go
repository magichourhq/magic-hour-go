package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiImageGeneratorCreateBody
type V1AiImageGeneratorCreateBody struct {
	// The aspect ratio of the output image(s). If not specified, defaults to `1:1` (square).
	AspectRatio nullable.Nullable[V1AiImageGeneratorCreateBodyAspectRatioEnum] `json:"aspect_ratio,omitempty"`
	// Number of images to generate. Maximum varies by model.
	ImageCount int `json:"image_count"`
	// The AI model to use for image generation. Each model has different capabilities and costs.
	//
	// **Models:**
	// - `default` - Use the model we recommend, which will change over time. This is recommended unless you need a specific model. This is the default behavior.
	// - `flux-schnell` - 5 credits/image
	//   - Supported resolutions: auto
	//   - Available for tiers: free, creator, pro, business
	//   - Image count allowed: 1, 2, 3, 4
	// - `z-image-turbo` - 5 credits/image
	//   - Supported resolutions: auto, 2k
	//   - Available for tiers: free, creator, pro, business
	//   - Image count allowed: 1, 2, 3, 4
	// - `seedream` - 30 credits/image
	//   - Supported resolutions: auto, 2k, 4k
	//   - Available for tiers: free, creator, pro, business
	//   - Image count allowed: 1, 2, 3, 4
	// - `nano-banana-pro` - 150 credits/image
	//   - Supported resolutions: auto
	//   - Available for tiers: creator, pro, business
	//   - Image count allowed: 1, 4, 9, 16
	//
	Model nullable.Nullable[V1AiImageGeneratorCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// DEPRECATED: Use `aspect_ratio` instead.
	//
	// The orientation of the output image(s). `aspect_ratio` takes precedence when `orientation` if both are provided.
	Orientation nullable.Nullable[V1AiImageGeneratorCreateBodyOrientationEnum] `json:"orientation,omitempty"`
	// Maximum resolution for the generated image.
	//
	// **Options:**
	// - `auto` - Automatic resolution (all tiers, default)
	// - `2k` - Up to 2048px (requires Pro or Business tier)
	// - `4k` - Up to 4096px (requires Business tier)
	//
	// Note: Resolution availability depends on the model and your subscription tier. See `model` field for which resolutions each model supports. Defaults to `auto` if not specified.
	Resolution nullable.Nullable[V1AiImageGeneratorCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// The art style to use for image generation.
	Style V1AiImageGeneratorCreateBodyStyle `json:"style"`
}
