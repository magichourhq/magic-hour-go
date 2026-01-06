package types

// Choose how to swap faces:
// **all-faces** (recommended) — swap all detected faces using one source image (`source_file_path` required)
// +- **individual-faces** — specify exact mappings using `face_mappings`
type V1FaceSwapCreateBodyAssetsFaceSwapModeEnum string

const (
	V1FaceSwapCreateBodyAssetsFaceSwapModeEnumAllFaces        V1FaceSwapCreateBodyAssetsFaceSwapModeEnum = "all-faces"
	V1FaceSwapCreateBodyAssetsFaceSwapModeEnumIndividualFaces V1FaceSwapCreateBodyAssetsFaceSwapModeEnum = "individual-faces"
)
