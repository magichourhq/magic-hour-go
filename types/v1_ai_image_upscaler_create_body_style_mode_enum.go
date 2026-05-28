package types

// The upscaling mode. `"pro"` is faster and does not require `enhancement`. `"creative"` requires `enhancement`. Defaults to `"creative"`.
type V1AiImageUpscalerCreateBodyStyleModeEnum string

const (
	V1AiImageUpscalerCreateBodyStyleModeEnumCreative V1AiImageUpscalerCreateBodyStyleModeEnum = "creative"
	V1AiImageUpscalerCreateBodyStyleModeEnumPro      V1AiImageUpscalerCreateBodyStyleModeEnum = "pro"
)
