package audio

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// ListRequest
type ListRequest struct {
	// The create endpoint to price, named by its path. For example `ai-voice-generator` for `POST /v1/ai-voice-generator`.
	Tool types.V1EstimateAudioListToolEnum `json:"tool"`
	// The number of characters in the `style.prompt` of the create request. Used by `ai-voice-generator`, `ai-voice-cloner`.
	CharacterCount nullable.Nullable[int] `json:"character_count,omitempty"`
	// The `duration` of the create request. Used by `ai-music-generator`.
	Duration nullable.Nullable[int] `json:"duration,omitempty"`
}
