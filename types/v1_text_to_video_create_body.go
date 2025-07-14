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
	// Controls the output video resolution. Defaults to `720p` if not specified.
	//
	// **Options:**
	// - `480p` - Supports only 5 or 10 second videos. Output: 24fps. Cost: 120 credits per 5 seconds.
	// - `720p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 300 credits per 5 seconds.
	// - `1080p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 600 credits per 5 seconds. **Requires** `pro` or `business` tier.
	Resolution nullable.Nullable[V1TextToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	Style      V1TextToVideoCreateBodyStyle                             `json:"style"`
}
