package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1ImageToVideoCreateBody
type V1ImageToVideoCreateBody struct {
	// Provide the assets for image-to-video.
	Assets V1ImageToVideoCreateBodyAssets `json:"assets"`
	// The total duration of the output video in seconds.
	EndSeconds float64 `json:"end_seconds"`
	// This field does not affect the output video's resolution. The video's orientation will match that of the input image.
	//
	// It is retained solely for backward compatibility and will be deprecated in the future.
	Height nullable.Nullable[int] `json:"height,omitempty"`
	// The name of video
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Controls the output video resolution. Defaults to `720p` if not specified.
	//
	// **Options:**
	// - `480p` - Supports only 5 or 10 second videos. Output: 24fps. Cost: 120 credits per 5 seconds.
	// - `720p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 300 credits per 5 seconds.
	// - `1080p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 600 credits per 5 seconds. **Requires** `pro` or `business` tier.
	Resolution nullable.Nullable[V1ImageToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Attributed used to dictate the style of the output
	Style nullable.Nullable[V1ImageToVideoCreateBodyStyle] `json:"style,omitempty"`
	// This field does not affect the output video's resolution. The video's orientation will match that of the input image.
	//
	// It is retained solely for backward compatibility and will be deprecated in the future.
	Width nullable.Nullable[int] `json:"width,omitempty"`
}
