package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiVoiceGeneratorCreateBody
type V1AiVoiceGeneratorCreateBody struct {
	// The name of audio. This value is mainly used for your own identification of the audio.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The content used to generate speech.
	Style V1AiVoiceGeneratorCreateBodyStyle `json:"style"`
}
