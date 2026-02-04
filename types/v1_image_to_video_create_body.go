package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1ImageToVideoCreateBody
type V1ImageToVideoCreateBody struct {
	// Provide the assets for image-to-video. Sora 2 only supports images with an aspect ratio of `9:16` or `16:9`.
	Assets V1ImageToVideoCreateBodyAssets `json:"assets"`
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
	// `height` is deprecated and no longer influences the output video's resolution.
	//
	// Output resolution is determined by the **minimum** of:
	// - The resolution of the input video
	// - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.
	//
	// This field is retained only for backward compatibility and will be removed in a future release.
	Height nullable.Nullable[int] `json:"height,omitempty"`
	// The AI model to use for video generation.
	// * `default`: Our recommended model for general use (Kling 2.5 Audio). Note: For backward compatibility, if you use default and end_seconds > 10, we'll fall back to Kling 1.6.
	// * `seedance`: Great for fast iteration and start/end frame
	// * `kling-2.5-audio`: Great for motion, action, and camera control
	// * `sora-2`: Great for story-telling, dialogue & creativity
	// * `veo3.1-audio`: Great for dialogue + SFX generated natively
	// * `veo3.1`: Great for realism, polish, & prompt adherence
	// * `kling-1.6`: Great for dependable clips with smooth motion
	Model nullable.Nullable[V1ImageToVideoCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Controls the output video resolution. Defaults to `720p` if not specified.
	//
	// * **Default**: Supports `480p`, `720p`, and `1080p`.
	// * **Seedance**: Supports `480p`, `720p`, `1080p`.
	// * **Kling 2.5 Audio**: Supports `720p`, `1080p`.
	// * **Sora 2**: Supports `720p`.
	// * **Veo 3.1 Audio**: Supports `720p`, `1080p`.
	// * **Veo 3.1**: Supports `720p`, `1080p`.
	// * **Kling 1.6**: Supports `720p`, `1080p`.
	Resolution nullable.Nullable[V1ImageToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	// Attributed used to dictate the style of the output
	Style nullable.Nullable[V1ImageToVideoCreateBodyStyle] `json:"style,omitempty"`
	// `width` is deprecated and no longer influences the output video's resolution.
	//
	// Output resolution is determined by the **minimum** of:
	// - The resolution of the input video
	// - The maximum resolution allowed by your subscription tier. See our [pricing page](https://magichour.ai/pricing) for more details.
	//
	// This field is retained only for backward compatibility and will be removed in a future release.
	Width nullable.Nullable[int] `json:"width,omitempty"`
}
