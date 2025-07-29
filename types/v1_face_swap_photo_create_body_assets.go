package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for face swap photo
type V1FaceSwapPhotoCreateBodyAssets struct {
	// This is the array of face mappings used for multiple face swap. The value is required if `face_swap_mode` is `individual-faces`.
	FaceMappings nullable.Nullable[[]V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem] `json:"face_mappings,omitempty"`
	// The mode of face swap.
	// * `all-faces` - Swap all faces in the target image or video. `source_file_path` is required.
	// * `individual-faces` - Swap individual faces in the target image or video. `source_faces` is required.
	FaceSwapMode nullable.Nullable[V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnum] `json:"face_swap_mode,omitempty"`
	// This is the image from which the face is extracted. The value is required if `face_swap_mode` is `all-faces`.
	//
	// This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	SourceFilePath nullable.Nullable[string] `json:"source_file_path,omitempty"`
	// This is the image where the face from the source image will be placed. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	TargetFilePath string `json:"target_file_path"`
}
