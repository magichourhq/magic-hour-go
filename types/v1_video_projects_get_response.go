package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Success
type V1VideoProjectsGetResponse struct {
	CreatedAt string `json:"created_at"`
	// The amount of credits deducted from your account to generate the video. If the status is not 'complete', this value is an estimate and may be adjusted upon completion based on the actual FPS of the output video.
	//
	// If video generation fails, credits will be refunded, and this field will be updated to include the refund.
	CreditsCharged int `json:"credits_charged"`
	// Deprecated: Please use `.downloads` instead. The download url and expiration date of the video project
	Download  nullable.Nullable[V1VideoProjectsGetResponseDownload] `json:"download,omitempty"`
	Downloads []V1VideoProjectsGetResponseDownloadsItem             `json:"downloads"`
	// Whether this resource is active. If false, it is deleted.
	Enabled bool `json:"enabled"`
	// End time of your clip (seconds). Must be greater than start_seconds.
	EndSeconds float64 `json:"end_seconds"`
	// In the case of an error, this object will contain the error encountered during video render
	Error nullable.Nullable[V1VideoProjectsGetResponseError] `json:"error,omitempty"`
	// Frame rate of the video. If the status is not 'complete', the frame rate is an estimate and will be adjusted when the video completes.
	Fps float64 `json:"fps"`
	// The height of the final output video. A value of -1 indicates the height can be ignored.
	Height int `json:"height"`
	// Unique ID of the video. Use it with the [Get video Project API](https://docs.magichour.ai/api-reference/video-projects/get-video-details) to fetch status and downloads.
	Id string `json:"id"`
	// The name of the video.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Start time of your clip (seconds). Must be ≥ 0.
	StartSeconds float64 `json:"start_seconds"`
	// The status of the video.
	Status V1VideoProjectsGetResponseStatusEnum `json:"status"`
	// Deprecated: Previously represented the number of frames (original name of our credit system) used for video generation. Use 'credits_charged' instead.
	//
	// The amount of frames used to generate the video. If the status is not 'complete', the cost is an estimate and will be adjusted when the video completes.
	TotalFrameCost int `json:"total_frame_cost"`
	// The type of the video project. Possible values are ANIMATION, IMAGE_TO_VIDEO, VIDEO_TO_VIDEO, TEXT_TO_VIDEO, FACE_SWAP, LIP_SYNC, AUTO_SUBTITLE, TALKING_PHOTO, UGC_AD, VIDEO_UPSCALER
	Type string `json:"type"`
	// The width of the final output video. A value of -1 indicates the width can be ignored.
	Width int `json:"width"`
}
