package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiMemeGeneratorCreateBody
type V1AiMemeGeneratorCreateBody struct {
	// The name of the meme.
	Name  nullable.Nullable[string]        `json:"name,omitempty"`
	Style V1AiMemeGeneratorCreateBodyStyle `json:"style"`
}
