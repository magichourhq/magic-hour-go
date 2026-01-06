package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiImageEditorCreateBody
type V1AiImageEditorCreateBody struct {
	// Provide the assets for image edit
	Assets V1AiImageEditorCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name  nullable.Nullable[string]      `json:"name,omitempty"`
	Style V1AiImageEditorCreateBodyStyle `json:"style"`
}
