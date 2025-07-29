package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for face swap. For video, The `video_source` field determines whether `video_file_path` or `youtube_url` field is used
type V1FaceSwapCreateBodyAssets struct {
	// This is the array of face mappings used for multiple face swap. The value is required if `face_swap_mode` is `individual-faces`.
	FaceMappings nullable.Nullable[[]V1FaceSwapCreateBodyAssetsFaceMappingsItem] `json:"face_mappings,omitempty"`
	// The mode of face swap.
	// * `all-faces` - Swap all faces in the target image or video. `source_file_path` is required.
	// * `individual-faces` - Swap individual faces in the target image or video. `source_faces` is required.
	FaceSwapMode nullable.Nullable[V1FaceSwapCreateBodyAssetsFaceSwapModeEnum] `json:"face_swap_mode,omitempty"`
	// The path of the input image with the face to be swapped.  The value is required if `face_swap_mode` is `all-faces`.
	//
	// This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	ImageFilePath nullable.Nullable[string] `json:"image_file_path,omitempty"`
	// The path of the input video. This field is required if `video_source` is `file`. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	VideoFilePath nullable.Nullable[string]                 `json:"video_file_path,omitempty"`
	VideoSource   V1FaceSwapCreateBodyAssetsVideoSourceEnum `json:"video_source"`
	// Using a youtube video as the input source. This field is required if `video_source` is `youtube`
	YoutubeUrl nullable.Nullable[string] `json:"youtube_url,omitempty"`
}
