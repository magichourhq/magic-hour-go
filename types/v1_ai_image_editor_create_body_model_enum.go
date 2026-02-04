package types

// The AI model to use for image editing. Each model has different capabilities and costs.
//
// **Models:**
// - `default` - Use the model we recommend, which will change over time. This is recommended unless you need a specific model. This is the default behavior.
// - `qwen-edit` - 10 credits/image
//   - Available for tiers: free, creator, pro, business
//   - Image count allowed: 1
//   - Max additional input images: 2
//
// - `nano-banana` - 50 credits/image
//   - Available for tiers: free, creator, pro, business
//   - Image count allowed: 1
//   - Max additional input images: 9
//
// - `seedream-v4` - 50 credits/image
//   - Available for tiers: free, creator, pro, business
//   - Image count allowed: 1
//   - Max additional input images: 9
//
// - `nano-banana-pro` - 150 credits/image
//   - Available for tiers: creator, pro, business
//   - Image count allowed: 1, 4, 9, 16
//   - Max additional input images: 9
//
// - `seedream-v4.5` - 100 credits/image
//   - Available for tiers: creator, pro, business
//   - Image count allowed: 1
//   - Max additional input images: 9
type V1AiImageEditorCreateBodyModelEnum string

const (
	V1AiImageEditorCreateBodyModelEnumDefault       V1AiImageEditorCreateBodyModelEnum = "default"
	V1AiImageEditorCreateBodyModelEnumNanoBanana    V1AiImageEditorCreateBodyModelEnum = "nano-banana"
	V1AiImageEditorCreateBodyModelEnumNanoBananaPro V1AiImageEditorCreateBodyModelEnum = "nano-banana-pro"
	V1AiImageEditorCreateBodyModelEnumQwenEdit      V1AiImageEditorCreateBodyModelEnum = "qwen-edit"
	V1AiImageEditorCreateBodyModelEnumSeedreamV4    V1AiImageEditorCreateBodyModelEnum = "seedream-v4"
	V1AiImageEditorCreateBodyModelEnumSeedreamV45   V1AiImageEditorCreateBodyModelEnum = "seedream-v4.5"
)
