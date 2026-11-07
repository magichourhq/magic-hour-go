package types

// Why the face detection failed. Present only when `status` is `error`; `null` otherwise.
type V1FaceDetectionGetResponseError struct {
	// An error code to indicate why a failure happened.
	Code string `json:"code"`
	// Details on the reason why a failure happened.
	Message string `json:"message"`
}
