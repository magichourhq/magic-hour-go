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
	// The height of the input video. This value will help determine the final orientation of the output video. The output video resolution may not match the input.
	Height int `json:"height"`
	// The name of video
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Attributed used to dictate the style of the output
	Style V1ImageToVideoCreateBodyStyle `json:"style"`
	// The width of the input video. This value will help determine the final orientation of the output video. The output video resolution may not match the input.
	Width int `json:"width"`
}
