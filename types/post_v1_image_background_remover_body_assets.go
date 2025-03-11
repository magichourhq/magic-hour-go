package types

// Provide the assets for background removal
type PostV1ImageBackgroundRemoverBodyAssets struct {
	// The image used to generate the image. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	ImageFilePath string `json:"image_file_path"`
}
