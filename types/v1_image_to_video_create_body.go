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
	// Attributed used to dictate the style of the output
	Style V1ImageToVideoCreateBodyStyle `json:"style"`
	// This field does not affect the output video's resolution. The video's orientation will match that of the input image.
	//
	// It is retained solely for backward compatibility and will be deprecated in the future.
	Width nullable.Nullable[int] `json:"width,omitempty"`
}
