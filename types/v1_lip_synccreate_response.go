package types

// Success
type V1LipSynccreateResponse struct {
	// Estimated cost of the video in terms of number of frames needed to render the video. Frames will be adjusted when the video completes
	EstimatedFrameCost int `json:"estimated_frame_cost"`
	// Unique ID of the video. This value can be used in the [get video project API](/api/tag/video-projects/get/v1/video-projects/{id}) to fetch additional details such as status
	Id string `json:"id"`
}
