package types

// Provide the assets for clothes changer
type V1AiClothesChangerCreateBodyAssets struct {
	// The image of the outfit. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file..
	GarmentFilePath string                                            `json:"garment_file_path"`
	GarmentType     V1AiClothesChangerCreateBodyAssetsGarmentTypeEnum `json:"garment_type"`
	// The image with the person. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file..
	PersonFilePath string `json:"person_file_path"`
}
