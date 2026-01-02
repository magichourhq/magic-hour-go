package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiVoiceClonerCreateBody
type V1AiVoiceClonerCreateBody struct {
	// Provide the assets for voice cloning.
	Assets V1AiVoiceClonerCreateBodyAssets `json:"assets"`
	// The name of audio. This value is mainly used for your own identification of the audio.
	Name  nullable.Nullable[string]      `json:"name,omitempty"`
	Style V1AiVoiceClonerCreateBodyStyle `json:"style"`
}
