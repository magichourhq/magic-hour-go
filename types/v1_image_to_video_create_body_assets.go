package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for image-to-video. Sora 2 only supports images with an aspect ratio of `9:16` or `16:9`.
type V1ImageToVideoCreateBodyAssets struct {
	// The image to use as the last frame of the video.
	//
	// * **`ltx-2`**: Not supported
	// * **`wan-2.2`**: Not supported
	// * **`seedance`**: Supports 480p, 720p, 1080p.
	// * **`seedance-2.0`**: Supports 480p, 720p.
	// * **`kling-2.5`**: Supports 1080p.
	// * **`kling-3.0`**: Supports 1080p.
	// * **`sora-2`**: Not supported
	// * **`veo3.1`**: Not supported
	// * **`veo3.1-lite`**: Not supported
	//
	// Legacy models:
	// * **`kling-1.6`**: Not supported
	EndImageFilePath nullable.Nullable[string] `json:"end_image_file_path,omitempty"`
	// The path of the image file. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// See the [file upload guide](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) for details.
	//
	ImageFilePath string `json:"image_file_path"`
}
