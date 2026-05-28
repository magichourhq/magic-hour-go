package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Style settings for the upscale. Use `mode` to select between `"pro"` (faster, no enhancement required) and `"creative"` (defaults to `"Balanced"` enhancement). Defaults to `"creative"`.
type V1AiImageUpscalerCreateBodyStyle struct {
	Enhancement nullable.Nullable[V1AiImageUpscalerCreateBodyStyleEnhancementEnum] `json:"enhancement,omitempty"`
	// The upscaling mode. `"pro"` is faster and does not require `enhancement`. `"creative"` requires `enhancement`. Defaults to `"creative"`.
	Mode nullable.Nullable[V1AiImageUpscalerCreateBodyStyleModeEnum] `json:"mode,omitempty"`
	// A prompt to guide the final image. This value is ignored if `enhancement` is not Creative
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
