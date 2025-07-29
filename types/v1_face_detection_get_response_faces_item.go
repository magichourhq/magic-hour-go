package types

// V1FaceDetectionGetResponseFacesItem
type V1FaceDetectionGetResponseFacesItem struct {
	// The path to the face image. This should be used in face swap photo/video API calls as `.assets.face_mappings.original_face`
	Path string `json:"path"`
	// The url to the face image. This is used to render the image in your applications.
	Url string `json:"url"`
}
