package types

// Success
type V1FaceSwapCreateResponse struct {
	// Estimated cost of the video in terms of number of frames needed to render the video. Frames will be adjusted when the video completes
	EstimatedFrameCost int `json:"estimated_frame_cost"`
	// Unique ID of the image. This value can be used in the [get image project API](https://docs.magichour.ai/api-reference/image-projects/get-image-details) to fetch additional details such as status
	Id string `json:"id"`
}
