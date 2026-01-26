package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Attributes used to dictate the style of the output
type V1AiTalkingPhotoCreateBodyStyle struct {
	// Controls overall motion style.
	// * `realistic` - Maintains likeness well, high quality, and reliable.
	// * `prompted` - Slightly lower likeness; allows option to prompt scene.
	//
	// **Deprecated values (maintained for backward compatibility):**
	// * `pro` - Deprecated: use `realistic`
	// * `standard` - Deprecated: use `prompted`
	// * `stable` - Deprecated: use `realistic`
	// * `expressive` - Deprecated: use `prompted`
	GenerationMode nullable.Nullable[V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum] `json:"generation_mode,omitempty"`
	// Note: this value is only applicable when generation_mode is `expressive`. The value can include up to 2 decimal places.
	// * Lower values yield more stability but can suppress mouth movement.
	// * Higher values increase motion and expressiveness, with a higher risk of distortion.
	Intensity nullable.Nullable[float64] `json:"intensity,omitempty"`
	// A text prompt to guide the generation. Only applicable when generation_mode is `prompted`.
	// This field is ignored for other modes.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
