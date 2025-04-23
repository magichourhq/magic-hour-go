package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiGifGeneratorCreateBody
type V1AiGifGeneratorCreateBody struct {
	// The name of gif
	Name  nullable.Nullable[string]       `json:"name,omitempty"`
	Style V1AiGifGeneratorCreateBodyStyle `json:"style"`
}
