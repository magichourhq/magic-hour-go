package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1TextToVideoCreateBody
type V1TextToVideoCreateBody struct {
	// Determines the aspect ratio of the output video.
	// * **Seedance**: Supports `9:16`, `16:9`, `1:1`.
	// * **Kling 2.5 Audio**: Supports `9:16`, `16:9`, `1:1`.
	// * **Sora 2**: Supports `9:16`, `16:9`.
	// * **Veo 3.1 Audio**: Supports `9:16`, `16:9`.
	// * **Veo 3.1**: Supports `9:16`, `16:9`.
	// * **Kling 1.6**: Supports `9:16`, `16:9`, `1:1`.
	AspectRatio nullable.Nullable[V1TextToVideoCreateBodyAspectRatioEnum] `json:"aspect_ratio,omitempty"`
	// The total duration of the output video in seconds.
	//
	// Supported durations depend on the chosen model:
	// * **Default**: 5-60 seconds (2-12 seconds for 480p).
	// * **Seedance**: 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12
	// * **Kling 2.5 Audio**: 5, 10
	// * **Sora 2**: 4, 8, 12, 24, 36, 48, 60
	// * **Veo 3.1 Audio**: 4, 6, 8, 16, 24, 32, 40, 48, 56
	// * **Veo 3.1**: 4, 6, 8, 16, 24, 32, 40, 48, 56
	// * **Kling 1.6**: 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60
	EndSeconds float64 `json:"end_seconds"`
	// The AI model to use for video generation.
	// * `default`: Our recommended model for general use (Kling 2.5 Audio). Note: For backward compatibility, if you use default and end_seconds > 10, we'll fall back to Kling 1.6.
	// * `seedance`: Great for fast iteration and start/end frame
	// * `kling-2.5-audio`: Great for motion, action, and camera control
	// * `sora-2`: Great for story-telling, dialogue & creativity
	// * `veo3.1-audio`: Great for dialogue + SFX generated natively
	// * `veo3.1`: Great for realism, polish, & prompt adherence
	// * `kling-1.6`: Great for dependable clips with smooth motion
	Model nullable.Nullable[V1TextToVideoCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Deprecated. Use `aspect_ratio` instead.
	Orientation nullable.Nullable[V1TextToVideoCreateBodyOrientationEnum] `json:"orientation,omitempty"`
	// Controls the output video resolution. Defaults to `720p` if not specified.
	//
	// * **Default**: Supports `480p`, `720p`, and `1080p`.
	// * **Seedance**: Supports `480p`, `720p`, `1080p`.
	// * **Kling 2.5 Audio**: Supports `720p`, `1080p`.
	// * **Sora 2**: Supports `720p`.
	// * **Veo 3.1 Audio**: Supports `720p`, `1080p`.
	// * **Veo 3.1**: Supports `720p`, `1080p`.
	// * **Kling 1.6**: Supports `720p`, `1080p`.
	Resolution nullable.Nullable[V1TextToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	Style      V1TextToVideoCreateBodyStyle                             `json:"style"`
}
