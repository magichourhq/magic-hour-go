package types

// Choose how to swap faces:
// **all-faces** (recommended) — swap all detected faces using one source image (`source_file_path` required)
// +- **individual-faces** — specify exact mappings using `face_mappings`
type V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnum string

const (
	V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumAllFaces        V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnum = "all-faces"
	V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumIndividualFaces V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnum = "individual-faces"
)
