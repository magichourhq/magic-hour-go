package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1TextToVideoCreateBody
type V1TextToVideoCreateBody struct {
	// The total duration of the output video in seconds.
	EndSeconds float64 `json:"end_seconds"`
	// The name of video
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Determines the orientation of the output video
	Orientation V1TextToVideoCreateBodyOrientationEnum `json:"orientation"`
	Style       V1TextToVideoCreateBodyStyle           `json:"style"`
}
