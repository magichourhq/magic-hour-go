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
	// - `flux-schnell` - from 5 credits/image
	//   - Supported resolutions: 640px, 1k, 2k
	//   - Available for tiers: free, creator, pro, business
	//   - Image count allowed: 1, 2, 3, 4
	// - `flux-2-klein` - from 5 credits/image
	//   - Supported resolutions: 640px, 1k, 2k
	//   - Available for tiers: free, creator, pro, business
	//   - Image count allowed: 1
	// - `z-image-turbo` - from 5 credits/image
	//   - Supported resolutions: 640px, 1k, 2k
	//   - Available for tiers: free, creator, pro, business
	//   - Image count allowed: 1, 2, 3, 4
	// - `seedream-v4` - from 40 credits/image
	//   - Supported resolutions: 640px, 1k, 2k, 4k
	//   - Available for tiers: creator, pro, business
	//   - Image count allowed: 1, 2, 3, 4
	// - `nano-banana` - from 50 credits/image
	//   - Supported resolutions: 640px, 1k
	//   - Available for tiers: creator, pro, business
	//   - Image count allowed: 1, 2, 3, 4
	// - `nano-banana-2` - from 100 credits/image
	//   - Supported resolutions: 640px, 1k, 2k, 4k
	//   - Available for tiers: creator, pro, business
	//   - Image count allowed: 1, 4, 9, 16
	// - `nano-banana-pro` - from 150 credits/image
	//   - Supported resolutions: 1k, 2k, 4k
	//   - Available for tiers: creator, pro, business
	//   - Image count allowed: 1, 4, 9, 16
	// - `gpt-image-2` - from 50 credits/image
	//   - Supported resolutions: 640px, 1k, 2k, 4k
	//   - Available for tiers: creator, pro, business
	//   - Image count allowed: 1, 2, 3, 4
	//
	// **Deprecated Enum Values:**
	// - `seedream` - Use `seedream-v4` instead.
	//
	Model nullable.Nullable[V1AiImageGeneratorCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// DEPRECATED: Use `aspect_ratio` instead.
	//
	// The orientation of the output image(s). `aspect_ratio` takes precedence when `orientation` if both are provided.
	Orientation nullable.Nullable[V1AiImageGeneratorCreateBodyOrientationEnum] `json:"orientation,omitempty"`
	// Maximum resolution (longest edge) for the output image.
	//
	// **Options:**
	// - `640px` — up to 640px
	// - `1k` — up to 1024px
	// - `2k` — up to 2048px
	// - `4k` — up to 4096px
	// - `auto` — **Deprecated.** Mapped server-side from your subscription tier to the best matching resolution the model supports
	//
	// **Per-model support:**
	// - `flux-schnell` - 640px, 1k, 2k
	// - `flux-2-klein` - 640px, 1k, 2k
	// - `z-image-turbo` - 640px, 1k, 2k
	// - `seedream-v4` - 640px, 1k, 2k, 4k
	// - `nano-banana` - 640px, 1k
	// - `nano-banana-2` - 640px, 1k, 2k, 4k
	// - `nano-banana-pro` - 1k, 2k, 4k
	// - `gpt-image-2` - 640px, 1k, 2k, 4k
	//
	// Note: Resolution availability depends on the model and your subscription tier.
	Resolution nullable.Nullable[V1AiImageGeneratorCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// The art style to use for image generation.
	Style V1AiImageGeneratorCreateBodyStyle `json:"style"`
}
