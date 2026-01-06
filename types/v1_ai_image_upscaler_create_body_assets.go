package types

// Provide the assets for upscaling
type V1AiImageUpscalerCreateBodyAssets struct {
	// The image to upscale. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	ImageFilePath string `json:"image_file_path"`
}
