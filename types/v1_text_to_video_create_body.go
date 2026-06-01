package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1TextToVideoCreateBody
type V1TextToVideoCreateBody struct {
	// Determines the aspect ratio of the output video.
	//
	// * **`ltx-2.3`**: Supports 9:16, 16:9, 1:1.
	// * **`wan-2.2`**: Supports 9:16, 16:9, 1:1.
	// * **`kling-2.5`**: Supports 9:16, 16:9, 1:1.
	// * **`kling-3.0`**: Supports 9:16, 16:9, 1:1.
	// * **`veo3.1-lite`**: Supports 9:16, 16:9.
	// * **`veo3.1`**: Supports 9:16, 16:9.
	// * **`seedance`**: Supports 9:16, 16:9, 1:1.
	// * **`seedance-2.0`**: Supports 9:16, 16:9, 1:1.
	// * **`sora-2`**: Supports 9:16, 16:9.
	//
	AspectRatio nullable.Nullable[V1TextToVideoCreateBodyAspectRatioEnum] `json:"aspect_ratio,omitempty"`
	// Whether to include audio in the video. Defaults to `false` if not specified.
	//
	// Audio support varies by model:
	// * **`ltx-2.3`**: Toggle-able: no additional credits for audio
	// * **`wan-2.2`**: Not supported
	// * **`kling-2.5`**: Toggle-able: no additional credits for audio
	// * **`kling-3.0`**: Toggle-able: audio adds extra credits when enabled
	// * **`veo3.1-lite`**: Toggle-able: audio adds extra credits when enabled
	// * **`veo3.1`**: Toggle-able: audio adds extra credits when enabled
	// * **`seedance`**: Not supported
	// * **`seedance-2.0`**: Toggle-able: no additional credits for audio
	// * **`sora-2`**: Toggle-able: no additional credits for audio
	//
	Audio nullable.Nullable[bool] `json:"audio,omitempty"`
	// The total duration of the output video in seconds. Supported durations depend on the chosen model:
	//
	// * **`ltx-2.3`**: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25, 30
	// * **`wan-2.2`**: 3, 4, 5, 6, 7, 8, 9, 10, 15
	// * **`kling-2.5`**: 5, 10
	// * **`kling-3.0`**: 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// * **`veo3.1-lite`**: 8, 16, 24, 32, 40, 48, 56
	// * **`veo3.1`**: 4, 6, 8, 16, 24, 32, 40, 48, 56
	// * **`seedance`**: 4, 5, 6, 7, 8, 9, 10, 11, 12
	// * **`seedance-2.0`**: 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// * **`sora-2`**: 4, 8, 12, 24, 36, 48, 60
	//
	EndSeconds float64 `json:"end_seconds"`
	// The AI model to use for video generation.
	//
	// * `default`: uses our currently recommended model for general use. For paid tiers, defaults to `kling-3.0`. For free tiers, it defaults to `ltx-2.3`.
	// * `ltx-2.3`: Fast iteration with lip-sync & end frame
	// * `wan-2.2`: Fast, strong visuals with effects
	// * `kling-2.5`: Motion, action, and camera control
	// * `kling-3.0`: Cinematic, multi-scene storytelling
	// * `veo3.1-lite`: Fast, affordable, high-quality
	// * `veo3.1`: Realistic visuals and prompt adherence
	// * `seedance`: Fast iteration
	// * `seedance-2.0`: State-of-the-art quality and consistency
	// * `sora-2`: Story-first concepts and creativity
	//
	// If you specify the deprecated model value that includes the `-audio` suffix, this will be the same as included `audio` as `true`.
	Model nullable.Nullable[V1TextToVideoCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Deprecated. Use `aspect_ratio` instead.
	Orientation nullable.Nullable[V1TextToVideoCreateBodyOrientationEnum] `json:"orientation,omitempty"`
	// Controls the output video resolution. Defaults to `720p` on paid tiers and `480p` on free tiers.
	//
	// * **`ltx-2.3`**: Supports 480p, 720p, 1080p.
	// * **`wan-2.2`**: Supports 480p, 720p, 1080p.
	// * **`kling-2.5`**: Supports 720p, 1080p.
	// * **`kling-3.0`**: Supports 720p, 1080p, 4k.
	// * **`veo3.1-lite`**: Supports 720p, 1080p.
	// * **`veo3.1`**: Supports 720p, 1080p.
	// * **`seedance`**: Supports 480p, 720p, 1080p.
	// * **`seedance-2.0`**: Supports 480p, 720p.
	// * **`sora-2`**: Supports 720p.
	//
	Resolution nullable.Nullable[V1TextToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	Style      V1TextToVideoCreateBodyStyle                             `json:"style"`
}
