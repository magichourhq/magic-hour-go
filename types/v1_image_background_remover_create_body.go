package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1ImageBackgroundRemoverCreateBody
type V1ImageBackgroundRemoverCreateBody struct {
	// Provide the assets for background removal
	Assets V1ImageBackgroundRemoverCreateBodyAssets `json:"assets"`
	// The name of image
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
