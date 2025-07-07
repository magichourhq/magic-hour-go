package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiImageEditorCreateBody
type V1AiImageEditorCreateBody struct {
	// Provide the assets for image edit
	Assets V1AiImageEditorCreateBodyAssets `json:"assets"`
	// The name of image
	Name  nullable.Nullable[string]      `json:"name,omitempty"`
	Style V1AiImageEditorCreateBodyStyle `json:"style"`
}
