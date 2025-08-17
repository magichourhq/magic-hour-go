package image_to_video

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for image-to-video.
	Assets types.V1ImageToVideoCreateBodyAssets `json:"assets"`
	// The total duration of the output video in seconds.
	EndSeconds float64 `json:"end_seconds"`
	// `height` is deprecated and no longer influences the output video's resolution.
	//
	// Output resolution is determined by the **minimum** of:
	// - The resolution of the input video
	// - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.
	//
	// This field is retained only for backward compatibility and will be removed in a future release.
	Height nullable.Nullable[int] `json:"height,omitempty"`
	// The name of video. This value is mainly used for your own identification of the video.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Controls the output video resolution. Defaults to `720p` if not specified.
	//
	// 480p and 720p are available on Creator, Pro, or Business tiers. However, 1080p require Pro or Business tier.
	//
	// **Options:**
	// - `480p` - Supports only 5 or 10 second videos. Output: 24fps. Cost: 120 credits per 5 seconds.
	// - `720p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 300 credits per 5 seconds.
	// - `1080p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 600 credits per 5 seconds.
	Resolution nullable.Nullable[types.V1ImageToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Attributed used to dictate the style of the output
	Style nullable.Nullable[types.V1ImageToVideoCreateBodyStyle] `json:"style,omitempty"`
	// `width` is deprecated and no longer influences the output video's resolution.
	//
	// Output resolution is determined by the **minimum** of:
	// - The resolution of the input video
	// - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.
	//
	// This field is retained only for backward compatibility and will be removed in a future release.
	Width nullable.Nullable[int] `json:"width,omitempty"`
}
