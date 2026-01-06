package types

// Provide the assets for face detection
type V1FaceDetectionCreateBodyAssets struct {
	// This is the image or video where the face will be detected. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	TargetFilePath string `json:"target_file_path"`
}
