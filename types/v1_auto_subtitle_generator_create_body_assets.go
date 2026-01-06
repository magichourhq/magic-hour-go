package types

// Provide the assets for auto subtitle generator
type V1AutoSubtitleGeneratorCreateBodyAssets struct {
	// This is the video used to add subtitles. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	VideoFilePath string `json:"video_file_path"`
}
