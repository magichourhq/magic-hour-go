
package types

// Success
type PostV1ImageToVideoResponse struct {
    // Estimated cost of the video in terms of number of frames needed to render the video. Frames will be adjusted when the video completes
    EstimatedFrameCost int `json:"estimated_frame_cost"`
    // Unique ID of the video. This value can be used in the [get video project API](https://docs.magichour.ai/api-reference/video-projects/get-video-details) to fetch additional details such as status
    Id string `json:"id"`
}



