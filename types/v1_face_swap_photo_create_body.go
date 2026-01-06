package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1FaceSwapPhotoCreateBody
type V1FaceSwapPhotoCreateBody struct {
	// Provide the assets for face swap photo
	Assets V1FaceSwapPhotoCreateBodyAssets `json:"assets"`
	// Give your image a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
