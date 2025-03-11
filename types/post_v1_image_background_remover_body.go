
package types
import (
nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1ImageBackgroundRemoverBody
type PostV1ImageBackgroundRemoverBody struct {
    // Provide the assets for background removal
    Assets PostV1ImageBackgroundRemoverBodyAssets `json:"assets"`
    // The name of image
    Name nullable.Nullable[string] `json:"name,omitempty"`
}



