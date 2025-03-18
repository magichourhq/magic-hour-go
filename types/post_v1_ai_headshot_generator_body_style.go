package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1AiHeadshotGeneratorBodyStyle
type PostV1AiHeadshotGeneratorBodyStyle struct {
	// A prompt to guide the final image.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
