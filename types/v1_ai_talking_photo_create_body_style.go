package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Attributes used to dictate the style of the output
type V1AiTalkingPhotoCreateBodyStyle struct {
	// Controls overall motion style.
	// * `expressive` - More motion and facial expressiveness; may introduce visual artifacts.
	// * `stable` -  Reduced motion for cleaner output; may result in minimal animation.
	GenerationMode nullable.Nullable[V1AiTalkingPhotoCreateBodyStyleGenerationModeEnum] `json:"generation_mode,omitempty"`
	// Note: this value is only applicable when generation_mode is `expressive`. The value can include up to 2 decimal places.
	// * Lower values yield more stability but can suppress mouth movement.
	// * Higher values increase motion and expressiveness, with a higher risk of distortion.
	Intensity nullable.Nullable[float64] `json:"intensity,omitempty"`
}
