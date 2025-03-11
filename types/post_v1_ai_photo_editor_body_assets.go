package types

// Provide the assets for photo editor
type PostV1AiPhotoEditorBodyAssets struct {
	// The image used to generate the output. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	ImageFilePath string `json:"image_file_path"`
}
