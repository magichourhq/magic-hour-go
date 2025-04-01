package types

// Provide the assets for creating a talking photo
type V1AiTalkingPhotoCreateBodyAssets struct {
	// The audio file to sync with the image. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	AudioFilePath string `json:"audio_file_path"`
	// The source image to animate. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	ImageFilePath string `json:"image_file_path"`
}
