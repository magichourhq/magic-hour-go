package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for lip-sync. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
type V1LipSyncCreateBodyAssets struct {
	// The path of the audio file. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	AudioFilePath string `json:"audio_file_path"`
	// Your video file. Required if `video_source` is `file`. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	VideoFilePath nullable.Nullable[string] `json:"video_file_path,omitempty"`
	// Choose your video source.
	VideoSource V1LipSyncCreateBodyAssetsVideoSourceEnum `json:"video_source"`
	// YouTube URL (required if `video_source` is `youtube`).
	YoutubeUrl nullable.Nullable[string] `json:"youtube_url,omitempty"`
}
