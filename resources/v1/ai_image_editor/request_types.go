package ai_image_editor

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// The aspect ratio of the output image(s). If not specified, defaults to `auto`.
	AspectRatio nullable.Nullable[types.V1AiImageEditorCreateBodyAspectRatioEnum] `json:"aspect_ratio,omitempty"`
	// Provide the assets for image edit
	Assets types.V1AiImageEditorCreateBodyAssets `json:"assets"`
	// Number of images to generate. Maximum varies by model. Defaults to 1 if not specified.
	ImageCount nullable.Nullable[float64] `json:"image_count,omitempty"`
	// The AI model to use for image editing. Each model has different capabilities and costs.
	//
	// **Models:**
	// - `default` - Use the model we recommend, which will change over time. This is recommended unless you need a specific model. This is the default behavior.
	// - `qwen-edit` - 10 credits/image
	//   - Available for tiers: free, creator, pro, business
	//   - Image count allowed: 1
	//   - Max additional input images: 2
	// - `nano-banana` - 50 credits/image
	//   - Available for tiers: free, creator, pro, business
	//   - Image count allowed: 1
	//   - Max additional input images: 9
	// - `seedream-v4` - 50 credits/image
	//   - Available for tiers: free, creator, pro, business
	//   - Image count allowed: 1
	//   - Max additional input images: 9
	// - `nano-banana-pro` - 150 credits/image
	//   - Available for tiers: creator, pro, business
	//   - Image count allowed: 1, 4, 9, 16
	//   - Max additional input images: 9
	// - `seedream-v4.5` - 100 credits/image
	//   - Available for tiers: creator, pro, business
	//   - Image count allowed: 1
	//   - Max additional input images: 9
	//
	Model nullable.Nullable[types.V1AiImageEditorCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your image a custom name for easy identification.
	Name  nullable.Nullable[string]            `json:"name,omitempty"`
	Style types.V1AiImageEditorCreateBodyStyle `json:"style"`
}
