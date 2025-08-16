package types

// Provide the assets for auto subtitle generator
type V1AutoSubtitleGeneratorCreateBodyAssets struct {
	// This is the video used to add subtitles. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.
	//
	VideoFilePath string `json:"video_file_path"`
}
