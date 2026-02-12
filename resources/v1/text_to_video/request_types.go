package text_to_video

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// Determines the aspect ratio of the output video.
	// * **seedance**: Supports `9:16`, `16:9`, `1:1`.
	// * **kling-2.5**: Supports `9:16`, `16:9`, `1:1`.
	// * **kling-3.0**: Supports `9:16`, `16:9`, `1:1`.
	// * **sora-2**: Supports `9:16`, `16:9`.
	// * **veo3.1**: Supports `9:16`, `16:9`.
	// * **kling-1.6**: Supports `9:16`, `16:9`, `1:1`.
	AspectRatio nullable.Nullable[types.V1TextToVideoCreateBodyAspectRatioEnum] `json:"aspect_ratio,omitempty"`
	// Whether to include audio in the video. Defaults to `false` if not specified.
	//
	// Audio support varies by model:
	// * **seedance**: Not supported
	// * **kling-2.5**: Always included (cannot be disabled)
	// * **kling-3.0**: Toggle-able (can enable/disable)
	// * **sora-2**: Always included (cannot be disabled)
	// * **veo3.1**: Toggle-able (can enable/disable)
	// * **kling-1.6**: Not supported
	Audio nullable.Nullable[bool] `json:"audio,omitempty"`
	// The total duration of the output video in seconds.
	//
	// Supported durations depend on the chosen model:
	// * **Default**: 5-60 seconds (2-12 seconds for 480p).
	// * **seedance**: 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12
	// * **kling-2.5**: 5, 10
	// * **kling-3.0**: 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// * **sora-2**: 4, 8, 12, 24, 36, 48, 60
	// * **veo3.1**: 4, 6, 8, 16, 24, 32, 40, 48, 56
	// * **kling-1.6**: 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60
	EndSeconds float64 `json:"end_seconds"`
	// The AI model to use for video generation.
	// * `default`: Our recommended model for general use (Kling 2.5 Audio). Note: For backward compatibility, if you use default and end_seconds > 10, we'll fall back to Kling 1.6.
	// * `seedance`: Great for fast iteration and start/end frame
	// * `kling-2.5`: Great for motion, action, and camera control
	// * `kling-3.0`: Great for cinematic, multi-scene storytelling with control
	// * `sora-2`: Great for story-telling, dialogue & creativity
	// * `veo3.1`: Great for realism, polish, & prompt adherence
	// * `kling-1.6`: Great for dependable clips with smooth motion
	Model nullable.Nullable[types.V1TextToVideoCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Deprecated. Use `aspect_ratio` instead.
	Orientation nullable.Nullable[types.V1TextToVideoCreateBodyOrientationEnum] `json:"orientation,omitempty"`
	// Controls the output video resolution. Defaults to `720p` if not specified.
	//
	// * **Default**: Supports `480p`, `720p`, and `1080p`.
	// * **seedance**: Supports `480p`, `720p`, `1080p`.
	// * **kling-2.5**: Supports `720p`, `1080p`.
	// * **kling-3.0**: Supports `720p`, `1080p`.
	// * **sora-2**: Supports `720p`.
	// * **veo3.1**: Supports `720p`, `1080p`.
	// * **kling-1.6**: Supports `720p`, `1080p`.
	Resolution nullable.Nullable[types.V1TextToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	Style      types.V1TextToVideoCreateBodyStyle                             `json:"style"`
}
