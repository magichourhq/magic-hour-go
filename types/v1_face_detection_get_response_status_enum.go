package types

// The status of the detection.
type V1FaceDetectionGetResponseStatusEnum string

const (
	V1FaceDetectionGetResponseStatusEnumComplete  V1FaceDetectionGetResponseStatusEnum = "complete"
	V1FaceDetectionGetResponseStatusEnumError     V1FaceDetectionGetResponseStatusEnum = "error"
	V1FaceDetectionGetResponseStatusEnumQueued    V1FaceDetectionGetResponseStatusEnum = "queued"
	V1FaceDetectionGetResponseStatusEnumRendering V1FaceDetectionGetResponseStatusEnum = "rendering"
)
