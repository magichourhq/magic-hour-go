package audio_projects

// DeleteRequest
type DeleteRequest struct {
	// Unique ID of the audio project. This value is returned by all of the POST APIs that create an audio.
	Id string `json:"id"`
}

// GetRequest
type GetRequest struct {
	// Unique ID of the audio project. This value is returned by all of the POST APIs that create an audio.
	Id string `json:"id"`
}
