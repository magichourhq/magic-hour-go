package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1LipSyncCreateBody
type V1LipSyncCreateBody struct {
	// Provide the assets for lip-sync. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
	Assets V1LipSyncCreateBodyAssets `json:"assets"`
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
	// Defines the maximum FPS (frames per second) for the output video. If the input video's FPS is lower than this limit, the output video will retain the input FPS. This is useful for reducing unnecessary frame usage in scenarios where high FPS is not required.
	MaxFpsLimit nullable.Nullable[float64] `json:"max_fps_limit,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds float64 `json:"start_seconds"`
	// Attributes used to dictate the style of the output
	Style nullable.Nullable[V1LipSyncCreateBodyStyle] `json:"style,omitempty"`
	// `width` is deprecated and no longer influences the output video's resolution.
	//
	// Output resolution is determined by the **minimum** of:
	// - The resolution of the input video
	// - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.
	//
	// This field is retained only for backward compatibility and will be removed in a future release.
	Width nullable.Nullable[int] `json:"width,omitempty"`
}
