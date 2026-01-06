package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1PhotoColorizerCreateBody
type V1PhotoColorizerCreateBody struct {
	// Provide the assets for photo colorization
	Assets V1PhotoColorizerCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
