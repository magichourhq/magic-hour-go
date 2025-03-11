
package ai_photo_editor
import (
types "github.com/magichourhq/magic-hour-go/types"
nullable "github.com/magichourhq/magic-hour-go/nullable"
)
// CreateRequest
type CreateRequest struct {
    // Provide the assets for photo editor
    Assets types.PostV1AiPhotoEditorBodyAssets `json:"assets"`
    // The name of image
    Name nullable.Nullable[string] `json:"name,omitempty"`
    // The resolution of the final output image. The allowed value is based on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
    Resolution int `json:"resolution"`
    // Deprecated: Please use `.style.steps` instead. Number of iterations used to generate the output. Higher values improve quality and increase the strength of the prompt but increase processing time.
    Steps nullable.Nullable[int] `json:"steps,omitempty"`
    Style types.PostV1AiPhotoEditorBodyStyle `json:"style"`
}



