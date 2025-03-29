package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1FaceSwapPhotoCreateBody
type V1FaceSwapPhotoCreateBody struct {
	// Provide the assets for face swap photo
	Assets V1FaceSwapPhotoCreateBodyAssets `json:"assets"`
	// The name of image
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
