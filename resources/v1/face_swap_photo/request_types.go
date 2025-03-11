
package face_swap_photo
import (
types "github.com/magichourhq/magic-hour-go/types"
nullable "github.com/magichourhq/magic-hour-go/nullable"
)
// CreateRequest
type CreateRequest struct {
    // Provide the assets for face swap photo
    Assets types.PostV1FaceSwapPhotoBodyAssets `json:"assets"`
    // The name of image
    Name nullable.Nullable[string] `json:"name,omitempty"`
}



