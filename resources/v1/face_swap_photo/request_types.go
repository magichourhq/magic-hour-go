package face_swap_photo

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for face swap photo
	Assets types.V1FaceSwapPhotoCreateBodyAssets `json:"assets"`
	// The name of image. This value is mainly used for your own identification of the image.
	Name nullable.Nullable[string] `json:"name,omitempty"`
}
