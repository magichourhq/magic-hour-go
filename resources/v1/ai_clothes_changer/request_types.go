
package ai_clothes_changer
import (
types "github.com/magichourhq/magic-hour-go/types"
nullable "github.com/magichourhq/magic-hour-go/nullable"
)
// CreateRequest
type CreateRequest struct {
    // Provide the assets for clothes changer
    Assets types.PostV1AiClothesChangerBodyAssets `json:"assets"`
    // The name of image
    Name nullable.Nullable[string] `json:"name,omitempty"`
}



