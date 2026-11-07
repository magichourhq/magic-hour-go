package video_to_video

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for video-to-video. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
	Assets types.V1VideoToVideoCreateBodyAssets `json:"assets"`
	// End time of your clip (seconds). Must be greater than start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// Source video FPS for v3 frame limits and estimated credits. Defaults to 30; final charges settle to rendered frames.
	Fps nullable.Nullable[float64] `json:"fps,omitempty"`
	// Determines whether the resulting video will have the same frame per second as the original video, or half.
	// * `FULL` - the result video will have the same FPS as the input video
	// * `HALF` - the result video will have half the FPS as the input video
	FpsResolution nullable.Nullable[types.V1VideoToVideoCreateBodyFpsResolutionEnum] `json:"fps_resolution,omitempty"`
	// Source video height in integer pixels. Required and positive for v3; ignored for v1/v2.
	Height nullable.Nullable[int] `json:"height,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds float64                             `json:"start_seconds"`
	Style        types.V1VideoToVideoCreateBodyStyle `json:"style"`
	// Source video width in integer pixels. Required and positive for v3; ignored for v1/v2.
	Width nullable.Nullable[int] `json:"width,omitempty"`
}
