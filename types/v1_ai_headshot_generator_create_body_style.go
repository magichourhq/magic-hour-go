package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiHeadshotGeneratorCreateBodyStyle
type V1AiHeadshotGeneratorCreateBodyStyle struct {
	// A prompt to guide the final image.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
