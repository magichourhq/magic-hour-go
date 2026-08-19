package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1TextToVideoCreateBody
type V1TextToVideoCreateBody struct {
	// Determines the aspect ratio of the output video.
	//
	// * **`kling-2.6`**: Supports 9:16, 16:9, 1:1.
	// * **`kling-3.0`**: Supports 9:16, 16:9, 1:1.
	// * **`ltx-2.3`**: Supports 9:16, 16:9, 1:1.
	// * **`minimax-h3`**: Supports 16:9, 9:16, 1:1.
	// * **`seedance-1.5`**: Supports 9:16, 16:9, 1:1.
	// * **`seedance-2.0`**: Supports 9:16, 16:9, 1:1.
	// * **`seedance-2.0-mini`**: Supports 9:16, 16:9, 1:1.
	// * **`seedance-2.5`**: Supports 9:16, 16:9, 1:1.
	// * **`sora-2`**: Supports 9:16, 16:9.
	// * **`veo3.1`**: Supports 9:16, 16:9.
	// * **`veo3.1-lite`**: Supports 9:16, 16:9.
	// * **`wan-2.2`**: Supports 9:16, 16:9, 1:1.
	//
	AspectRatio nullable.Nullable[V1TextToVideoCreateBodyAspectRatioEnum] `json:"aspect_ratio,omitempty"`
	// Whether to include audio in the video. Defaults to `false` if not specified.
	//
	// Audio support varies by model:
	// * **`kling-2.6`**: Not supported
	// * **`kling-3.0`**: Toggle-able: audio adds extra credits when enabled
	// * **`ltx-2.3`**: Toggle-able: no additional credits for audio
	// * **`minimax-h3`**: Toggle-able: no additional credits for audio
	// * **`seedance-1.5`**: Toggle-able: audio adds extra credits when enabled
	// * **`seedance-2.0`**: Toggle-able: no additional credits for audio
	// * **`seedance-2.0-mini`**: Toggle-able: no additional credits for audio
	// * **`seedance-2.5`**: Toggle-able: no additional credits for audio
	// * **`sora-2`**: Toggle-able: no additional credits for audio
	// * **`veo3.1`**: Toggle-able: audio adds extra credits when enabled
	// * **`veo3.1-lite`**: Toggle-able: audio adds extra credits when enabled
	// * **`wan-2.2`**: Not supported
	//
	Audio nullable.Nullable[bool] `json:"audio,omitempty"`
	// The total duration of the output video in seconds. Supported durations depend on the chosen model:
	//
	// * **`kling-2.6`**: 5, 10
	// * **`kling-3.0`**: 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// * **`ltx-2.3`**: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25, 30
	// * **`minimax-h3`**: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 25, 30
	// * **`seedance-1.5`**: 4, 5, 6, 7, 8, 9, 10, 11, 12
	// * **`seedance-2.0`**: 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// * **`seedance-2.0-mini`**: 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15
	// * **`seedance-2.5`**: 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30
	// * **`sora-2`**: 4, 8, 12, 24, 36, 48, 60
	// * **`veo3.1`**: 4, 6, 8, 16, 24, 32, 40, 48, 56
	// * **`veo3.1-lite`**: 4, 6, 8, 16, 24, 32, 40, 48, 56
	// * **`wan-2.2`**: 3, 4, 5, 6, 7, 8, 9, 10, 15
	//
	EndSeconds float64 `json:"end_seconds"`
	// The AI model to use for video generation.
	//
	// * `default`: uses our currently recommended model for general use. For paid tiers, defaults to `kling-3.0`. For free tiers, it defaults to `ltx-2.3`.
	// * `kling-2.6`: Great for action, motion blur, and camera moves.
	// * `kling-3.0`: Best overall quality for cinematic storytelling.
	// * `ltx-2.3`: Fastest output. Best for rapid iteration.
	// * `minimax-h3`: Reference-driven video with native audio.
	// * `seedance-1.5`: Smooth, consistent motion with precision.
	// * `seedance-2.0`: Top quality with reference-to-video control.
	// * `seedance-2.0-mini`: Fast, consistent video with strong motion quality
	// * `seedance-2.5`: Highest quality with superior realism, detail, and motion
	// * `sora-2`: Open AI's model. Great for creativity and viral clips.
	// * `veo3.1`: Google's model. Highest realism and detail.
	// * `veo3.1-lite`: Veo quality at a more accessible cost.
	// * `wan-2.2`: Strong physics, camera moves, and motion.
	//
	// If you specify the deprecated model value that includes the `-audio` suffix, this will be the same as included `audio` as `true`.
	Model nullable.Nullable[V1TextToVideoCreateBodyModelEnum] `json:"model,omitempty"`
	// Give your video a custom name for easy identification.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Deprecated. Use `aspect_ratio` instead.
	Orientation nullable.Nullable[V1TextToVideoCreateBodyOrientationEnum] `json:"orientation,omitempty"`
	// Controls the output video resolution. Defaults to `720p` on paid tiers and `480p` on free tiers.
	//
	// * **`kling-2.6`**: Supports 720p, 1080p.
	// * **`kling-3.0`**: Supports 720p, 1080p, 4k.
	// * **`ltx-2.3`**: Supports 480p, 720p, 1080p.
	// * **`minimax-h3`**: Supports 480p, 720p, 1080p.
	// * **`seedance-1.5`**: Supports 480p, 720p, 1080p.
	// * **`seedance-2.0`**: Supports 480p, 720p.
	// * **`seedance-2.0-mini`**: Supports 480p, 720p.
	// * **`seedance-2.5`**: Supports 480p, 720p.
	// * **`sora-2`**: Supports 720p.
	// * **`veo3.1`**: Supports 720p, 1080p.
	// * **`veo3.1-lite`**: Supports 720p, 1080p.
	// * **`wan-2.2`**: Supports 480p, 720p, 1080p.
	//
	Resolution nullable.Nullable[V1TextToVideoCreateBodyResolutionEnum] `json:"resolution,omitempty"`
	Style      V1TextToVideoCreateBodyStyle                             `json:"style"`
}
