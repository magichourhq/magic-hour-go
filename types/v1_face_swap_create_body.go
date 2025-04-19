package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1FaceSwapCreateBody
type V1FaceSwapCreateBody struct {
	// Provide the assets for face swap. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
	Assets V1FaceSwapCreateBodyAssets `json:"assets"`
	// The end time of the input video in seconds
	EndSeconds float64 `json:"end_seconds"`
	// Used to determine the dimensions of the output video.
	//
	// * If height is provided, width will also be required. The larger value between width and height will be used to determine the maximum output resolution while maintaining the original aspect ratio.
	// * If both height and width are omitted, the video will be resized according to your subscription's maximum resolution, while preserving aspect ratio.
	//
	// Note: if the video's original resolution is less than the maximum, the video will not be resized.
	//
	// See our [pricing page](https://magichour.ai/pricing) for more details.
	Height nullable.Nullable[int] `json:"height,omitempty"`
	// The name of video
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The start time of the input video in seconds
	StartSeconds float64 `json:"start_seconds"`
	// Used to determine the dimensions of the output video.
	//
	// * If width is provided, height will also be required. The larger value between width and height will be used to determine the maximum output resolution while maintaining the original aspect ratio.
	// * If both height and width are omitted, the video will be resized according to your subscription's maximum resolution, while preserving aspect ratio.
	//
	// Note: if the video's original resolution is less than the maximum, the video will not be resized.
	//
	// See our [pricing page](https://magichour.ai/pricing) for more details.
	Width nullable.Nullable[int] `json:"width,omitempty"`
}
