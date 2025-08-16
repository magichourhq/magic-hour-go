package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for video-to-video. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
type V1VideoToVideoCreateBodyAssets struct {
	// Required if `video_source` is `file`. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.
	//
	VideoFilePath nullable.Nullable[string]                     `json:"video_file_path,omitempty"`
	VideoSource   V1VideoToVideoCreateBodyAssetsVideoSourceEnum `json:"video_source"`
	// Using a youtube video as the input source. This field is required if `video_source` is `youtube`
	YoutubeUrl nullable.Nullable[string] `json:"youtube_url,omitempty"`
}
