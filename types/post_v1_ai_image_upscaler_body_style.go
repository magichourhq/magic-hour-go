package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1AiImageUpscalerBodyStyle
type PostV1AiImageUpscalerBodyStyle struct {
	Enhancement PostV1AiImageUpscalerBodyStyleEnhancementEnum `json:"enhancement"`
	// A prompt to guide the final image. This value is ignored if `enhancement` is not Creative
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
