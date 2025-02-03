package types

// PostV1FilesUploadUrlsResponseItemsItem
type PostV1FilesUploadUrlsResponseItemsItem struct {
	// when the upload url expires, and will need to request a new one.
	ExpiresAt string `json:"expires_at"`
	// this value is used in APIs that needs assets, such as image_file_path, video_file_path, and audio_file_path
	FilePath string `json:"file_path"`
	// Used to upload the file to storage, send a PUT request with the file as data to upload.
	UploadUrl string `json:"upload_url"`
}
