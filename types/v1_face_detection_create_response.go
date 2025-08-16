package types

// V1FaceDetectionCreateResponse
type V1FaceDetectionCreateResponse struct {
	// The credits charged for the task.
	CreditsCharged int `json:"credits_charged"`
	// The id of the task. Use this value in the [get face detection details API](/api-reference/files/get-face-detection-details) to get the details of the face detection task.
	Id string `json:"id"`
}
