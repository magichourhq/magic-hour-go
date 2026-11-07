package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1FaceDetectionGetResponse
type V1FaceDetectionGetResponse struct {
	// The credits charged for the task.
	CreditsCharged int `json:"credits_charged"`
	// Why the face detection failed. Present only when `status` is `error`; `null` otherwise.
	Error nullable.Nullable[V1FaceDetectionGetResponseError] `json:"error,omitempty"`
	// The faces detected in the image or video. The list is populated as faces are detected.
	Faces []V1FaceDetectionGetResponseFacesItem `json:"faces"`
	// The id of the task. This value is returned by the [face detection API](https://docs.magichour.ai/api-reference/files/face-detection#response-id).
	Id string `json:"id"`
	// The status of the detection.
	Status V1FaceDetectionGetResponseStatusEnum `json:"status"`
}
