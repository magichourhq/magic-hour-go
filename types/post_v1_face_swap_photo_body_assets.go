package types

// Provide the assets for face swap photo
type PostV1FaceSwapPhotoBodyAssets struct {
	// This is the image from which the face is extracted. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	SourceFilePath string `json:"source_file_path"`
	// This is the image where the face from the source image will be placed. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	TargetFilePath string `json:"target_file_path"`
}
