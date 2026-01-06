package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for image edit
type V1AiImageEditorCreateBodyAssets struct {
	// Deprecated: Please use `image_file_paths` instead as edits with multiple images are now supported. The image used in the edit. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	ImageFilePath nullable.Nullable[string] `json:"image_file_path,omitempty"`
	// The image(s) used in the edit, maximum of 10 images. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	ImageFilePaths nullable.Nullable[[]string] `json:"image_file_paths,omitempty"`
}
