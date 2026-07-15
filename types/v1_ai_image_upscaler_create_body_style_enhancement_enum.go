package types

// Deprecated: use `mode` instead. `"Resemblance"` maps to `"preserve"`. `"Balanced"` and `"Creative"` map to the same-named modes.
type V1AiImageUpscalerCreateBodyStyleEnhancementEnum string

const (
	V1AiImageUpscalerCreateBodyStyleEnhancementEnumBalanced    V1AiImageUpscalerCreateBodyStyleEnhancementEnum = "Balanced"
	V1AiImageUpscalerCreateBodyStyleEnhancementEnumCreative    V1AiImageUpscalerCreateBodyStyleEnhancementEnum = "Creative"
	V1AiImageUpscalerCreateBodyStyleEnhancementEnumResemblance V1AiImageUpscalerCreateBodyStyleEnhancementEnum = "Resemblance"
)
