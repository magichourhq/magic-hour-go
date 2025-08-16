package types

// Provide the assets for clothes changer
type V1AiClothesChangerCreateBodyAssets struct {
	// The image of the outfit. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.
	//
	GarmentFilePath string `json:"garment_file_path"`
	// The type of the outfit.
	GarmentType V1AiClothesChangerCreateBodyAssetsGarmentTypeEnum `json:"garment_type"`
	// The image with the person. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.
	//
	PersonFilePath string `json:"person_file_path"`
}
