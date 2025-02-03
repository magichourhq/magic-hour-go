package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1AiHeadshotGeneratorBody
type PostV1AiHeadshotGeneratorBody struct {
	// Provide the assets for headshot photo
	Assets PostV1AiHeadshotGeneratorBodyAssets `json:"assets"`
	// The name of image
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
