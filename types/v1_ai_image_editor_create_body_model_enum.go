package types

// The AI model to use for image editing. Each model has different capabilities and costs.
//
// **Models:**
// - `default` - Use the model we recommend, which will change over time. This is recommended unless you need a specific model. This is the default behavior.
// - `qwen-edit` - from 10 credits/image
//   - Supported resolutions: 640px, 1k, 2k
//   - Available for tiers: free, creator, pro, business
//   - Max additional input images: 2
//
// - `nano-banana` - from 50 credits/image
//   - Supported resolutions: 640px, 1k
//   - Available for tiers: creator, pro, business
//   - Max additional input images: 9
//
// - `nano-banana-2` - from 100 credits/image
//   - Supported resolutions: 640px, 1k, 2k, 4k
//   - Available for tiers: creator, pro, business
//   - Max additional input images: 9
//
// - `seedream-v4` - from 40 credits/image
//   - Supported resolutions: 640px, 1k, 2k, 4k
//   - Available for tiers: creator, pro, business
//   - Max additional input images: 9
//
// - `nano-banana-pro` - from 150 credits/image
//   - Supported resolutions: 1k, 2k, 4k
//   - Available for tiers: creator, pro, business
//   - Max additional input images: 9
//
// - `seedream-v4.5` - from 50 credits/image
//   - Supported resolutions: 640px, 1k, 2k, 4k
//   - Available for tiers: creator, pro, business
//   - Max additional input images: 9
//
// - `gpt-image-2` - from 50 credits/image
//   - Supported resolutions: 640px, 1k, 2k, 4k
//   - Available for tiers: creator, pro, business
//   - Max additional input images: 9
type V1AiImageEditorCreateBodyModelEnum string

const (
	V1AiImageEditorCreateBodyModelEnumDefault       V1AiImageEditorCreateBodyModelEnum = "default"
	V1AiImageEditorCreateBodyModelEnumGptImage2     V1AiImageEditorCreateBodyModelEnum = "gpt-image-2"
	V1AiImageEditorCreateBodyModelEnumNanoBanana    V1AiImageEditorCreateBodyModelEnum = "nano-banana"
	V1AiImageEditorCreateBodyModelEnumNanoBanana2   V1AiImageEditorCreateBodyModelEnum = "nano-banana-2"
	V1AiImageEditorCreateBodyModelEnumNanoBananaPro V1AiImageEditorCreateBodyModelEnum = "nano-banana-pro"
	V1AiImageEditorCreateBodyModelEnumQwenEdit      V1AiImageEditorCreateBodyModelEnum = "qwen-edit"
	V1AiImageEditorCreateBodyModelEnumSeedreamV4    V1AiImageEditorCreateBodyModelEnum = "seedream-v4"
	V1AiImageEditorCreateBodyModelEnumSeedreamV45   V1AiImageEditorCreateBodyModelEnum = "seedream-v4.5"
)
