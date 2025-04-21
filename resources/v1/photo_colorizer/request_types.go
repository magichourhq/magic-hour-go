package photo_colorizer

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for photo colorization
	Assets types.V1PhotoColorizerCreateBodyAssets `json:"assets"`
	// The name of image
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
