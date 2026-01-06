package ai_image_editor

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for image edit
	Assets types.V1AiImageEditorCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name  nullable.Nullable[string]            `json:"name,omitempty"`
	Style types.V1AiImageEditorCreateBodyStyle `json:"style"`
}
