
package types
import (
nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1AnimationBody
type PostV1AnimationBody struct {
    // Provide the assets for animation.
    Assets PostV1AnimationBodyAssets `json:"assets"`
    // The end time of the input video in seconds
    EndSeconds float64 `json:"end_seconds"`
    // The desire output video frame rate
    Fps float64 `json:"fps"`
    // The height of the final output video. The maximum height depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
    Height int `json:"height"`
    // The name of video
    Name nullable.Nullable[string] `json:"name,omitempty"`
    // Defines the style of the output video
    Style PostV1AnimationBodyStyle `json:"style"`
    // The width of the final output video. The maximum width depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
    Width int `json:"width"`
}



