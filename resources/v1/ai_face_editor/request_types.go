package ai_face_editor

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for face editor
	Assets types.V1AiFaceEditorCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Face editing parameters
	Style types.V1AiFaceEditorCreateBodyStyle `json:"style"`
}
