package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiClothesChangerCreateBody
type V1AiClothesChangerCreateBody struct {
	// Provide the assets for clothes changer
	Assets V1AiClothesChangerCreateBodyAssets `json:"assets"`
	// The name of image. This value is mainly used for your own identification of the image.
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
