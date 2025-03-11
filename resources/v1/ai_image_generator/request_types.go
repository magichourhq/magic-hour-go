
package ai_image_generator
import (
types "github.com/magichourhq/magic-hour-go/types"
nullable "github.com/magichourhq/magic-hour-go/nullable"
)
// CreateRequest
type CreateRequest struct {
    // number to images to generate
    ImageCount int `json:"image_count"`
    // The name of image
    Name nullable.Nullable[string] `json:"name,omitempty"`
    Orientation types.PostV1AiImageGeneratorBodyOrientationEnum `json:"orientation"`
    Style types.PostV1AiImageGeneratorBodyStyle `json:"style"`
}



