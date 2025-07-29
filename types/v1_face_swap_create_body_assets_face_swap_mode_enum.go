package types

// The mode of face swap.
// * `all-faces` - Swap all faces in the target image or video. `source_file_path` is required.
// * `individual-faces` - Swap individual faces in the target image or video. `source_faces` is required.
type V1FaceSwapCreateBodyAssetsFaceSwapModeEnum string

const (
	V1FaceSwapCreateBodyAssetsFaceSwapModeEnumAllFaces        V1FaceSwapCreateBodyAssetsFaceSwapModeEnum = "all-faces"
	V1FaceSwapCreateBodyAssetsFaceSwapModeEnumIndividualFaces V1FaceSwapCreateBodyAssetsFaceSwapModeEnum = "individual-faces"
)
