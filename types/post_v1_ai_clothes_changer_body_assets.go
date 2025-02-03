package types

// Provide the assets for clothes changer
type PostV1AiClothesChangerBodyAssets struct {
	// The image of the outfit. This value can be the `file_path` field from the response of the [upload urls API](/docs/api/tag/files/post/v1/files/upload-urls), or the url of the file.
	GarmentFilePath string                                          `json:"garment_file_path"`
	GarmentType     PostV1AiClothesChangerBodyAssetsGarmentTypeEnum `json:"garment_type"`
	// The image with the person. This value can be the `file_path` field from the response of the [upload urls API](/docs/api/tag/files/post/v1/files/upload-urls), or the url of the file.
	PersonFilePath string `json:"person_file_path"`
}
