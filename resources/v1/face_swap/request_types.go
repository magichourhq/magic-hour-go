package face_swap

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for face swap. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
	Assets types.V1FaceSwapCreateBodyAssets `json:"assets"`
	// End time of your clip (seconds). Must be greater than start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// `height` is deprecated and no longer influences the output video's resolution.
	//
	// Output resolution is determined by the **minimum** of:
	// - The resolution of the input video
	// - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.
	//
	// This field is retained only for backward compatibility and will be removed in a future release.
	Height nullable.Nullable[int] `json:"height,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds float64 `json:"start_seconds"`
	// Style of the face swap video.
	Style nullable.Nullable[types.V1FaceSwapCreateBodyStyle] `json:"style,omitempty"`
	// `width` is deprecated and no longer influences the output video's resolution.
	//
	// Output resolution is determined by the **minimum** of:
	// - The resolution of the input video
	// - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.
	//
	// This field is retained only for backward compatibility and will be removed in a future release.
	Width nullable.Nullable[int] `json:"width,omitempty"`
}
