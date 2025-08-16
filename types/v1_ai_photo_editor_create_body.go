package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiPhotoEditorCreateBody
type V1AiPhotoEditorCreateBody struct {
	// Provide the assets for photo editor
	Assets V1AiPhotoEditorCreateBodyAssets `json:"assets"`
	// The name of image. This value is mainly used for your own identification of the image.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The resolution of the final output image. The allowed value is based on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
	Resolution int `json:"resolution"`
	// Deprecated: Please use `.style.steps` instead. Number of iterations used to generate the output. Higher values improve quality and increase the strength of the prompt but increase processing time.
	Steps nullable.Nullable[int]         `json:"steps,omitempty"`
	Style V1AiPhotoEditorCreateBodyStyle `json:"style"`
}
