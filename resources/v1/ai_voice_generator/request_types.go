package ai_voice_generator

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// The name of audio. This value is mainly used for your own identification of the audio.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The content used to generate speech.
	Style types.V1AiVoiceGeneratorCreateBodyStyle `json:"style"`
}
