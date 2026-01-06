package animation

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Provide the assets for animation.
	Assets types.V1AnimationCreateBodyAssets `json:"assets"`
	// This value determines the duration of the output video.
	EndSeconds float64 `json:"end_seconds"`
	// The desire output video frame rate
	Fps float64 `json:"fps"`
	// The height of the final output video. The maximum height depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
	Height int `json:"height"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Defines the style of the output video
	Style types.V1AnimationCreateBodyStyle `json:"style"`
	// The width of the final output video. The maximum width depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
	Width int `json:"width"`
}
