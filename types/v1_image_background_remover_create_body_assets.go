package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for background removal
type V1ImageBackgroundRemoverCreateBodyAssets struct {
	// The image used as the new background for the image_file_path. This image will be resized to match the image in image_file_path. Please make sure the resolution between the images are similar.
	//
	// This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	BackgroundImageFilePath nullable.Nullable[string] `json:"background_image_file_path,omitempty"`
	// The image to remove the background. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
	ImageFilePath string `json:"image_file_path"`
}
