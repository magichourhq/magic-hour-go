// Code generated by Sideko (sideko.dev) DO NOT EDIT.

package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1VideoToVideoBody
type PostV1VideoToVideoBody struct {
	// Provide the assets for video-to-video. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
	Assets PostV1VideoToVideoBodyAssets `json:"assets"`
	// The end time of the input video in seconds
	EndSeconds float64 `json:"end_seconds"`
	// Determines whether the resulting video will have the same frame per second as the original video, or half.
	// * `FULL` - the result video will have the same FPS as the input video
	// * `HALF` - the result video will have half the FPS as the input video
	FpsResolution nullable.Nullable[PostV1VideoToVideoBodyFpsResolutionEnum] `json:"fps_resolution,omitempty"`
	// The height of the final output video. Must be divisible by 64. The maximum height depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
	Height float64 `json:"height"`
	// The name of video
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// The start time of the input video in seconds
	StartSeconds float64                     `json:"start_seconds"`
	Style        PostV1VideoToVideoBodyStyle `json:"style"`
	// The width of the final output video. Must be divisible by 64. The maximum width depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
	Width float64 `json:"width"`
}
