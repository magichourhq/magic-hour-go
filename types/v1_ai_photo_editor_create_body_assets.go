package types

// Provide the assets for photo editor
type V1AiPhotoEditorCreateBodyAssets struct {
	// The image used to generate the output. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.
	//
	ImageFilePath string `json:"image_file_path"`
}
