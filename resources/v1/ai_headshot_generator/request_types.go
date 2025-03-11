
package ai_headshot_generator
import (
types "github.com/magichourhq/magic-hour-go/types"
nullable "github.com/magichourhq/magic-hour-go/nullable"
)
// CreateRequest
type CreateRequest struct {
    // Provide the assets for headshot photo
    Assets types.PostV1AiHeadshotGeneratorBodyAssets `json:"assets"`
    // The name of image
    Name nullable.Nullable[string] `json:"name,omitempty"`
}



