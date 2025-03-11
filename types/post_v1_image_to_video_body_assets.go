package types

// Provide the assets for image-to-video.
type PostV1ImageToVideoBodyAssets struct {
	// The path of the image file. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	ImageFilePath string `json:"image_file_path"`
}
