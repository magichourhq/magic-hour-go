package text_to_video

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// The total duration of the output video in seconds.
	EndSeconds float64 `json:"end_seconds"`
	// The name of video
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Determines the orientation of the output video
	Orientation types.V1TextToVideoCreateBodyOrientationEnum `json:"orientation"`
	Style       types.V1TextToVideoCreateBodyStyle           `json:"style"`
}
