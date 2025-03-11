
package types
import (
nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1FaceSwapBody
type PostV1FaceSwapBody struct {
    // Provide the assets for face swap. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
    Assets PostV1FaceSwapBodyAssets `json:"assets"`
    // The end time of the input video in seconds
    EndSeconds float64 `json:"end_seconds"`
    // The height of the final output video. The maximum height depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
    Height int `json:"height"`
    // The name of video
    Name nullable.Nullable[string] `json:"name,omitempty"`
    // The start time of the input video in seconds
    StartSeconds float64 `json:"start_seconds"`
    // The width of the final output video. The maximum width depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
    Width int `json:"width"`
}



