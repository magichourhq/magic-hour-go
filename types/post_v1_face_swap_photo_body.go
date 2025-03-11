
package types
import (
nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1FaceSwapPhotoBody
type PostV1FaceSwapPhotoBody struct {
    // Provide the assets for face swap photo
    Assets PostV1FaceSwapPhotoBodyAssets `json:"assets"`
    // The name of image
    Name nullable.Nullable[string] `json:"name,omitempty"`
}



