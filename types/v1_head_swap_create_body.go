package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1HeadSwapCreateBody
type V1HeadSwapCreateBody struct {
	// Provide the body and head images for head swap
	Assets V1HeadSwapCreateBodyAssets `json:"assets"`
	// Constrains the larger dimension (height or width) of the output. Omit to use the maximum allowed for your plan (capped at 2048px). Values above your plan maximum are clamped down to your plan's maximum.
	MaxResolution nullable.Nullable[int] `json:"max_resolution,omitempty"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
