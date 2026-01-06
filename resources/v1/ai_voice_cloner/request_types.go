package ai_voice_cloner

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for voice cloning.
	Assets types.V1AiVoiceClonerCreateBodyAssets `json:"assets"`
	// Give your audio a custom name for easy identification.
	Name  nullable.Nullable[string]            `json:"name,omitempty"`
	Style types.V1AiVoiceClonerCreateBodyStyle `json:"style"`
}
