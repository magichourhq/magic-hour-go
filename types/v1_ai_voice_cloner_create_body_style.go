package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiVoiceClonerCreateBodyStyle
type V1AiVoiceClonerCreateBodyStyle struct {
	// Spoken language. Use a canonical lowercase language name, or auto to detect the language from the text. Defaults to auto.
	Language nullable.Nullable[V1AiVoiceClonerCreateBodyStyleLanguageEnum] `json:"language,omitempty"`
	// Text used to generate speech from the cloned voice. The character limit is 5000 characters.
	Prompt string `json:"prompt"`
	// Speech speed multiplier, from 0.75 to 1.30 inclusive. Defaults to 1 (normal speed).
	Speed nullable.Nullable[float64] `json:"speed,omitempty"`
}
