package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for clothes changer
type V1AiClothesChangerCreateBodyAssets struct {
	// The image of the outfit. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	GarmentFilePath string `json:"garment_file_path"`
	// Deprecated: garment_type is no longer needed.
	GarmentType nullable.Nullable[V1AiClothesChangerCreateBodyAssetsGarmentTypeEnum] `json:"garment_type,omitempty"`
	// The image with the person. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	PersonFilePath string `json:"person_file_path"`
}
