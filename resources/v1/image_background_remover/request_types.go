
package image_background_remover
import (
types "github.com/magichourhq/magic-hour-go/types"
nullable "github.com/magichourhq/magic-hour-go/nullable"
)
// CreateRequest
type CreateRequest struct {
    // Provide the assets for background removal
    Assets types.PostV1ImageBackgroundRemoverBodyAssets `json:"assets"`
    // The name of image
    Name nullable.Nullable[string] `json:"name,omitempty"`
}



