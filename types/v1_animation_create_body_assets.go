package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for animation.
type V1AnimationCreateBodyAssets struct {
	// The path of the input audio. This field is required if `audio_source` is `file`. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.
	//
	AudioFilePath nullable.Nullable[string] `json:"audio_file_path,omitempty"`
	// Optionally add an audio source if you'd like to incorporate audio into your video
	AudioSource V1AnimationCreateBodyAssetsAudioSourceEnum `json:"audio_source"`
	// An initial image to use a the first frame of the video. This value is either
	// - a direct URL to the video file
	// - `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls).
	//
	// Please refer to the [Input File documentation](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls#input-file) to learn more.
	//
	ImageFilePath nullable.Nullable[string] `json:"image_file_path,omitempty"`
	// Using a youtube video as the input source. This field is required if `audio_source` is `youtube`
	YoutubeUrl nullable.Nullable[string] `json:"youtube_url,omitempty"`
}
