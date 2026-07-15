package types

// The upscaling mode. `"preserve"` uses the fast pro pipeline (1× credit multiplier). `"balanced"` and `"creative"` use the creative pipeline (2× credit multiplier). `"pro"` is deprecated and maps to `"preserve"`. Defaults to `"balanced"`.
type V1AiImageUpscalerCreateBodyStyleModeEnum string

const (
	V1AiImageUpscalerCreateBodyStyleModeEnumBalanced V1AiImageUpscalerCreateBodyStyleModeEnum = "balanced"
	V1AiImageUpscalerCreateBodyStyleModeEnumCreative V1AiImageUpscalerCreateBodyStyleModeEnum = "creative"
	V1AiImageUpscalerCreateBodyStyleModeEnumPreserve V1AiImageUpscalerCreateBodyStyleModeEnum = "preserve"
	V1AiImageUpscalerCreateBodyStyleModeEnumPro      V1AiImageUpscalerCreateBodyStyleModeEnum = "pro"
)
