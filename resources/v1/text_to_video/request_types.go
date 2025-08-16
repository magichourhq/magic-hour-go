package text_to_video

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// The total duration of the output video in seconds.
	//
	// The value must be greater than or equal to 5 seconds and less than or equal to 60 seconds.
	//
	// Note: For 480p resolution, the value must be either 5 or 10.
	EndSeconds float64 `json:"end_seconds"`
	// The name of video. This value is mainly used for your own identification of the video.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Determines the orientation of the output video
	Orientation types.V1TextToVideoCreateBodyOrientationEnum `json:"orientation"`
	// Controls the output video resolution. Defaults to `720p` if not specified.
	//
	// 480p and 720p are available on Creator, Pro, or Business tiers. However, 1080p require Pro or Business tier.
	//
	// **Options:**
	// - `480p` - Supports only 5 or 10 second videos. Output: 24fps. Cost: 120 credits per 5 seconds.
	// - `720p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 300 credits per 5 seconds.
	// - `1080p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 600 credits per 5 seconds.
	Resolution nullable.Nullable[types.V1TextToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	Style      types.V1TextToVideoCreateBodyStyle                             `json:"style"`
}
