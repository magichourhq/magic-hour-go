package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Attributes used to dictate the style of the output
type V1AudioToVideoCreateBodyStyle struct {
	// Prompt to guide the visual style of the video.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
