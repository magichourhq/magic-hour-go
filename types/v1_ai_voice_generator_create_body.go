package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiVoiceGeneratorCreateBody
type V1AiVoiceGeneratorCreateBody struct {
	// Give your audio a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The content used to generate speech.
	Style V1AiVoiceGeneratorCreateBodyStyle `json:"style"`
}
