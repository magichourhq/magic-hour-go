package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for face swap. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
type V1FaceSwapCreateBodyAssets struct {
	// This is the array of face mappings used for multiple face swap. The value is required if `face_swap_mode` is `individual-faces`.
	FaceMappings nullable.Nullable[[]V1FaceSwapCreateBodyAssetsFaceMappingsItem] `json:"face_mappings,omitempty"`
	// Choose how to swap faces:
	// **all-faces** (recommended) — swap all detected faces using one source image (`source_file_path` required)
	// +- **individual-faces** — specify exact mappings using `face_mappings`
	FaceSwapMode nullable.Nullable[V1FaceSwapCreateBodyAssetsFaceSwapModeEnum] `json:"face_swap_mode,omitempty"`
	// The path of the input image with the face to be swapped.  The value is required if `face_swap_mode` is `all-faces`.
	//
	// This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	ImageFilePath nullable.Nullable[string] `json:"image_file_path,omitempty"`
	// Your video file. Required if `video_source` is `file`. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	VideoFilePath nullable.Nullable[string] `json:"video_file_path,omitempty"`
	// Choose your video source.
	VideoSource V1FaceSwapCreateBodyAssetsVideoSourceEnum `json:"video_source"`
	// YouTube URL (required if `video_source` is `youtube`).
	YoutubeUrl nullable.Nullable[string] `json:"youtube_url,omitempty"`
}
