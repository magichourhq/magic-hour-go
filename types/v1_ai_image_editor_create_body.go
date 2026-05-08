package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiImageEditorCreateBody
type V1AiImageEditorCreateBody struct {
	// The aspect ratio of the output image(s). If not specified, defaults to `auto`.
	AspectRatio nullable.Nullable[V1AiImageEditorCreateBodyAspectRatioEnum] `json:"aspect_ratio,omitempty"`
	// Provide the assets for image edit
	Assets V1AiImageEditorCreateBodyAssets `json:"assets"`
	// Number of images to generate. Maximum varies by model. Defaults to 1 if not specified.
	ImageCount nullable.Nullable[float64] `json:"image_count,omitempty"`
	// The AI model to use for image editing. Each model has different capabilities and costs.
	//
	// **Models:**
	// - `default` - Use the model we recommend, which will change over time. This is recommended unless you need a specific model. This is the default behavior.
	// - `qwen-edit` - from 10 credits/image
	//   - Supported resolutions: 640px, 1k, 2k
	//   - Available for tiers: free, creator, pro, business
	//   - Max additional input images: 2
	// - `nano-banana` - from 50 credits/image
	//   - Supported resolutions: 640px, 1k
	//   - Available for tiers: free, creator, pro, business
	//   - Max additional input images: 9
	// - `nano-banana-2` - from 100 credits/image
	//   - Supported resolutions: 640px, 1k, 2k, 4k
	//   - Available for tiers: free, creator, pro, business
	//   - Max additional input images: 9
	// - `seedream-v4` - from 40 credits/image
	//   - Supported resolutions: 640px, 1k, 2k, 4k
	//   - Available for tiers: free, creator, pro, business
	//   - Max additional input images: 9
	// - `nano-banana-pro` - from 150 credits/image
	//   - Supported resolutions: 1k, 2k, 4k
	//   - Available for tiers: creator, pro, business
	//   - Max additional input images: 9
	// - `seedream-v4.5` - from 50 credits/image
	//   - Supported resolutions: 640px, 1k, 2k, 4k
	//   - Available for tiers: creator, pro, business
	//   - Max additional input images: 9
	// - `gpt-image-2` - from 50 credits/image
	//   - Supported resolutions: 640px, 1k, 2k, 4k
	//   - Available for tiers: creator, pro, business
	//   - Max additional input images: 9
	//
	Model nullable.Nullable[V1AiImageEditorCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
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
	// - `qwen-edit` - 640px, 1k, 2k
	// - `nano-banana` - 640px, 1k
	// - `nano-banana-2` - 640px, 1k, 2k, 4k
	// - `seedream-v4` - 640px, 1k, 2k, 4k
	// - `nano-banana-pro` - 1k, 2k, 4k
	// - `seedream-v4.5` - 640px, 1k, 2k, 4k
	// - `gpt-image-2` - 640px, 1k, 2k, 4k
	//
	// Note: Resolution availability depends on the model and your subscription tier.
	Resolution nullable.Nullable[V1AiImageEditorCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	Style      V1AiImageEditorCreateBodyStyle                             `json:"style"`
}
