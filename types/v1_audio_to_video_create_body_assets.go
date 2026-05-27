package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the audio file and an optional reference image.
type V1AudioToVideoCreateBodyAssets struct {
	// The path of the audio file. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	AudioFilePath string `json:"audio_file_path"`
	// Reference image for the initial frame of the video. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	ImageFilePath nullable.Nullable[string] `json:"image_file_path,omitempty"`
}
