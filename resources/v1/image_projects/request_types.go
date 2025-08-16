package image_projects

// DeleteRequest
type DeleteRequest struct {
	// Unique ID of the image project. This value is returned by all of the POST APIs that create an image.
	Id string `json:"id"`
}

// GetRequest
type GetRequest struct {
	// Unique ID of the image project. This value is returned by all of the POST APIs that create an image.
	Id string `json:"id"`
}
