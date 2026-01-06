package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1ImageBackgroundRemoverCreateBody
type V1ImageBackgroundRemoverCreateBody struct {
	// Provide the assets for background removal
	Assets V1ImageBackgroundRemoverCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
