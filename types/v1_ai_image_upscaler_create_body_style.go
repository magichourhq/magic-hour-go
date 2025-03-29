package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiImageUpscalerCreateBodyStyle
type V1AiImageUpscalerCreateBodyStyle struct {
	Enhancement V1AiImageUpscalerCreateBodyStyleEnhancementEnum `json:"enhancement"`
	// A prompt to guide the final image. This value is ignored if `enhancement` is not Creative
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
