package types

// Provide the assets for auto subtitle generator
type V1AutoSubtitleGeneratorCreateBodyAssets struct {
	// This is the video used to add subtitles. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	VideoFilePath string `json:"video_file_path"`
}
