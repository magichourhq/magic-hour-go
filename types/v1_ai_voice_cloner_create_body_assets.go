package types

// Provide the assets for voice cloning.
type V1AiVoiceClonerCreateBodyAssets struct {
	// The audio used to clone the voice. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.
	//
	AudioFilePath string `json:"audio_file_path"`
}
