package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1AiClothesChangerBody
type PostV1AiClothesChangerBody struct {
	// Provide the assets for clothes changer
	Assets PostV1AiClothesChangerBodyAssets `json:"assets"`
	// The name of image
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
