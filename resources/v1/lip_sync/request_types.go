package lip_sync

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for lip-sync. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
	Assets types.V1LipSyncCreateBodyAssets `json:"assets"`
	// The end time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0.1, and more than the start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// `height` is deprecated and no longer influences the output video's resolution.
	//
	// Output resolution is determined by the **minimum** of:
	// - The resolution of the input video
	// - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.
	//
	// This field is retained only for backward compatibility and will be removed in a future release.
	Height nullable.Nullable[int] `json:"height,omitempty"`
	// Defines the maximum FPS (frames per second) for the output video. If the input video's FPS is lower than this limit, the output video will retain the input FPS. This is useful for reducing unnecessary frame usage in scenarios where high FPS is not required.
	MaxFpsLimit nullable.Nullable[float64] `json:"max_fps_limit,omitempty"`
	// The name of video. This value is mainly used for your own identification of the video.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The start time of the input video in seconds. This value is used to trim the input video. The value must be greater than 0.
	StartSeconds float64 `json:"start_seconds"`
	// `width` is deprecated and no longer influences the output video's resolution.
	//
	// Output resolution is determined by the **minimum** of:
	// - The resolution of the input video
	// - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.
	//
	// This field is retained only for backward compatibility and will be removed in a future release.
	Width nullable.Nullable[int] `json:"width,omitempty"`
}
