// Code generated by Sideko (sideko.dev) DO NOT EDIT.

package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Success
type GetV1VideoProjectsIdResponse struct {
	CreatedAt string `json:"created_at"`
	// The download url and expiration date of the video project
	Download nullable.Nullable[GetV1VideoProjectsIdResponseDownload] `json:"download"`
	// The end time of the input video in seconds
	EndSeconds float64 `json:"end_seconds"`
	// Frame rate of the video. If the status is not 'complete', the frame rate is an estimate and will be adjusted when the video completes.
	Fps float64 `json:"fps"`
	// The height of the final output video. The maximum height depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
	Height float64 `json:"height"`
	// Unique ID of the video. This value can be used in the [get video project API](/api/tag/video-projects/get/v1/video-projects/{id}) to fetch additional details such as status
	Id string `json:"id"`
	// The name of the video.
	Name nullable.Nullable[string] `json:"name"`
	// The start time of the input video in seconds
	StartSeconds float64 `json:"start_seconds"`
	// The status of the video.
	Status GetV1VideoProjectsIdResponseStatusEnum `json:"status"`
	// The amount of frames used to generate the video. If the status is not 'complete', the cost is an estimate and will be adjusted when the video completes.
	TotalFrameCost float64                              `json:"total_frame_cost"`
	Type           GetV1VideoProjectsIdResponseTypeEnum `json:"type"`
	// The width of the final output video. The maximum width depends on your subscription. Please refer to our [pricing page](https://magichour.ai/pricing) for more details
	Width float64 `json:"width"`
}
