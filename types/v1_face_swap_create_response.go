package types

// Success
type V1FaceSwapCreateResponse struct {
	// The amount of credits deducted from your account to generate the video. If the status is not 'complete', this value is an estimate and may be adjusted upon completion based on the actual FPS of the output video.
	//
	// If video generation fails, credits will be refunded, and this field will be updated to include the refund.
	CreditsCharged int `json:"credits_charged"`
	// Deprecated: Previously represented the number of frames (original name of our credit system) used for video generation. Use 'credits_charged' instead.
	//
	// The amount of frames used to generate the video. If the status is not 'complete', the cost is an estimate and will be adjusted when the video completes.
	EstimatedFrameCost int `json:"estimated_frame_cost"`
	// Unique ID of the video. Use it with the [Get video Project API](https://docs.magichour.ai/api-reference/video-projects/get-video-details) to fetch status and downloads.
	Id string `json:"id"`
}
