package lip_sync

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for lip-sync. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
	Assets types.V1LipSyncCreateBodyAssets `json:"assets"`
	// The end time of the input video in seconds
	EndSeconds float64 `json:"end_seconds"`
	// The height of the final output video. The maximum height depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
	Height int `json:"height"`
	// Defines the maximum FPS (frames per second) for the output video. If the input video's FPS is lower than this limit, the output video will retain the input FPS. This is useful for reducing unnecessary frame usage in scenarios where high FPS is not required.
	MaxFpsLimit nullable.Nullable[float64] `json:"max_fps_limit,omitempty"`
	// The name of video
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The start time of the input video in seconds
	StartSeconds float64 `json:"start_seconds"`
	// The width of the final output video. The maximum width depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
	Width int `json:"width"`
}
