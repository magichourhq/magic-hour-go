package ai_voice_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Give your audio a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The content used to generate speech.
	Style types.V1AiVoiceGeneratorCreateBodyStyle `json:"style"`
}
