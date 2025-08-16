package video_projects

// DeleteRequest
type DeleteRequest struct {
	// Unique ID of the video project. This value is returned by all of the POST APIs that create a video.
	Id string `json:"id"`
}

// GetRequest
type GetRequest struct {
	// Unique ID of the video project. This value is returned by all of the POST APIs that create a video.
	Id string `json:"id"`
}
