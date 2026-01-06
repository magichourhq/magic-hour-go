package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for face swap photo
type V1FaceSwapPhotoCreateBodyAssets struct {
	// This is the array of face mappings used for multiple face swap. The value is required if `face_swap_mode` is `individual-faces`.
	FaceMappings nullable.Nullable[[]V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem] `json:"face_mappings,omitempty"`
	// Choose how to swap faces:
	// **all-faces** (recommended) — swap all detected faces using one source image (`source_file_path` required)
	// +- **individual-faces** — specify exact mappings using `face_mappings`
	FaceSwapMode nullable.Nullable[V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnum] `json:"face_swap_mode,omitempty"`
	// This is the image from which the face is extracted. The value is required if `face_swap_mode` is `all-faces`.
	//
	// This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	SourceFilePath nullable.Nullable[string] `json:"source_file_path,omitempty"`
	// This is the image where the face from the source image will be placed. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	TargetFilePath string `json:"target_file_path"`
}
