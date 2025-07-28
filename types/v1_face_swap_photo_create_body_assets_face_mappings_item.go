package types

// V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem
type V1FaceSwapPhotoCreateBodyAssetsFaceMappingsItem struct {
	// The face image that will be used to replace the face in the `original_face`. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	NewFace string `json:"new_face"`
	// The face detected from the image in `target_file_path`. The file name is in the format of `<face_frame>-<face_index>.png`. This value is corresponds to the response in the face detection API [COMING SOON].
	//
	// * The face_frame is the frame number of the face in the target image. For images, the frame number is always 0.
	// * The face_index is the index of the face in the target image, starting from 0 going left to right.
	OriginalFace string `json:"original_face"`
}
