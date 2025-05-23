package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1AiPhotoEditorCreateBodyStyle
type V1AiPhotoEditorCreateBodyStyle struct {
	// Use this to describe what your input image is. This helps maintain aspects of the image you don't want to change.
	ImageDescription string `json:"image_description"`
	// Determines the input image's influence. Higher values align the output more with the initial image.
	LikenessStrength float64 `json:"likeness_strength"`
	// What you want to avoid seeing in the final output; has a minor effect.
	NegativePrompt nullable.Nullable[string] `json:"negative_prompt,omitempty"`
	// What you want your final output to look like. We recommend starting with the image description and making minor edits for best results.
	Prompt string `json:"prompt"`
	// Determines the prompt's influence. Higher values align the output more with the prompt.
	PromptStrength float64 `json:"prompt_strength"`
	// Number of iterations used to generate the output. Higher values improve quality and increase the strength of the prompt but increase processing time.
	Steps nullable.Nullable[int] `json:"steps,omitempty"`
	// The multiplier applied to an image's original dimensions during the upscaling process. For example, a scale of 2 doubles the width and height (e.g., from 512x512 to 1024x1024).
	UpscaleFactor nullable.Nullable[int] `json:"upscale_factor,omitempty"`
	// Upscale fidelity refers to the level of quality desired in the generated image. Fidelity value of 1 means more details.
	UpscaleFidelity nullable.Nullable[float64] `json:"upscale_fidelity,omitempty"`
}
