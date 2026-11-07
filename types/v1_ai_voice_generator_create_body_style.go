package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// The content used to generate speech.
type V1AiVoiceGeneratorCreateBodyStyle struct {
	// Spoken language. Use a canonical lowercase language name, or auto to detect the language from the text. Defaults to auto.
	Language nullable.Nullable[V1AiVoiceGeneratorCreateBodyStyleLanguageEnum] `json:"language,omitempty"`
	// Text used to generate speech. The character limit is 5000 characters.
	Prompt string `json:"prompt"`
	// Speech speed multiplier, from 0.75 to 1.30 inclusive. Defaults to 1 (normal speed).
	Speed nullable.Nullable[float64] `json:"speed,omitempty"`
	// The voice to use for the speech. Use the `id` of a voice returned by `GET /v1/voices`.
	VoiceName V1AiVoiceGeneratorCreateBodyStyleVoiceNameEnum `json:"voice_name"`
}
