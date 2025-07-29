package types

// The mode of face swap.
// * `all-faces` - Swap all faces in the target image or video. `source_file_path` is required.
// * `individual-faces` - Swap individual faces in the target image or video. `source_faces` is required.
type V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnum string

const (
	V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumAllFaces        V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnum = "all-faces"
	V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnumIndividualFaces V1FaceSwapPhotoCreateBodyAssetsFaceSwapModeEnum = "individual-faces"
)
