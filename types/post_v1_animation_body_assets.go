
package types
import (
nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Provide the assets for animation.
type PostV1AnimationBodyAssets struct {
    // The path of the input audio. This field is required if `audio_source` is `file`. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
    AudioFilePath nullable.Nullable[string] `json:"audio_file_path,omitempty"`
    // Optionally add an audio source if you'd like to incorporate audio into your video
    AudioSource PostV1AnimationBodyAssetsAudioSourceEnum `json:"audio_source"`
    // An initial image to use a the first frame of the video. This value can be either the `file_path` field from the response of the [upload urls API](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls), or the url of the file.
    ImageFilePath nullable.Nullable[string] `json:"image_file_path,omitempty"`
    // Using a youtube video as the input source. This field is required if `audio_source` is `youtube`
    YoutubeUrl nullable.Nullable[string] `json:"youtube_url,omitempty"`
}



