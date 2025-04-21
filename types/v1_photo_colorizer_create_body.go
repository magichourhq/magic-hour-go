package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1PhotoColorizerCreateBody
type V1PhotoColorizerCreateBody struct {
	// Provide the assets for photo colorization
	Assets V1PhotoColorizerCreateBodyAssets `json:"assets"`
	// The name of image
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
