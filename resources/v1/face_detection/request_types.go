package face_detection

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// GetRequest
type GetRequest struct {
	// The id of the task
	Id string `json:"id"`
}

// CreateRequest
type CreateRequest struct {
	// Provide the assets for face detection
	Assets types.V1FaceDetectionCreateBodyAssets `json:"assets"`
	// Confidence threshold for filtering detected faces.
	// * Higher values (e.g., 0.9) include only faces detected with high certainty, reducing false positives.
	// * Lower values (e.g., 0.3) include more faces, but may increase the chance of incorrect detections.
	ConfidenceScore nullable.Nullable[float64] `json:"confidence_score,omitempty"`
}
