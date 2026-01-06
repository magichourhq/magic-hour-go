package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiFaceEditorCreateBody
type V1AiFaceEditorCreateBody struct {
	// Provide the assets for face editor
	Assets V1AiFaceEditorCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Face editing parameters
	Style V1AiFaceEditorCreateBodyStyle `json:"style"`
}
