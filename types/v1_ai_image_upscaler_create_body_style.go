package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Style settings for the upscale. Use `mode` (`"preserve"`, `"balanced"`, or `"creative"`). Defaults to `"balanced"`.
type V1AiImageUpscalerCreateBodyStyle struct {
	// Deprecated: use `mode` instead. `"Resemblance"` maps to `"preserve"`. `"Balanced"` and `"Creative"` map to the same-named modes.
	Enhancement nullable.Nullable[V1AiImageUpscalerCreateBodyStyleEnhancementEnum] `json:"enhancement,omitempty"`
	// The upscaling mode. `"preserve"` uses the fast pro pipeline (1× credit multiplier). `"balanced"` and `"creative"` use the creative pipeline (2× credit multiplier). `"pro"` is deprecated and maps to `"preserve"`. Defaults to `"balanced"`.
	Mode nullable.Nullable[V1AiImageUpscalerCreateBodyStyleModeEnum] `json:"mode,omitempty"`
	// A prompt to guide the final image. Only used when mode is `creative`.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
