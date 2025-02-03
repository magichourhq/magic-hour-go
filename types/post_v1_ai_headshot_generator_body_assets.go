package types

// Provide the assets for headshot photo
type PostV1AiHeadshotGeneratorBodyAssets struct {
	// The image used to generate the headshot. This image must contain one detectable face. This is the `file_path` field from the response of the [upload urls API](/docs/api/tag/files/post/v1/files/upload-urls)
	ImageFilePath string `json:"image_file_path"`
}
