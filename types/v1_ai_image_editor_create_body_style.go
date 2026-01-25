package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiImageEditorCreateBodyStyle
type V1AiImageEditorCreateBodyStyle struct {
	// The AI model to use for image editing.
	// * `Nano Banana` - Precise, realistic edits with consistent results
	// * `Seedream` - Creative, imaginative images with artistic freedom
	// * `default` - Use the model we recommend, which will change over time. This is recommended unless you need a specific model. This is the default behavior.
	Model nullable.Nullable[V1AiImageEditorCreateBodyStyleModelEnum] `json:"model,omitempty"`
	// The prompt used to edit the image.
	Prompt string `json:"prompt"`
}
