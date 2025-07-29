package types

// V1FaceDetectionGetResponse
type V1FaceDetectionGetResponse struct {
	// The credits charged for the task.
	CreditsCharged int `json:"credits_charged"`
	// The faces detected in the image or video. The list is populated as faces are detected.
	Faces []V1FaceDetectionGetResponseFacesItem `json:"faces"`
	// The id of the task
	Id string `json:"id"`
	// The status of the detection.
	Status V1FaceDetectionGetResponseStatusEnum `json:"status"`
}
