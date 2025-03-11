package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for face swap. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
type PostV1FaceSwapBodyAssets struct {
	// The path of the input image. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	ImageFilePath string `json:"image_file_path"`
	// The path of the input video. This field is required if `video_source` is `file`. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	VideoFilePath nullable.Nullable[string]               `json:"video_file_path,omitempty"`
	VideoSource   PostV1FaceSwapBodyAssetsVideoSourceEnum `json:"video_source"`
	// Using a youtube video as the input source. This field is required if `video_source` is `youtube`
	YoutubeUrl nullable.Nullable[string] `json:"youtube_url,omitempty"`
}
