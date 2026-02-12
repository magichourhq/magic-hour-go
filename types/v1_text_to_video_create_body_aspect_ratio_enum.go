package types

// Determines the aspect ratio of the output video.
// * **seedance**: Supports `9:16`, `16:9`, `1:1`.
// * **kling-2.5**: Supports `9:16`, `16:9`, `1:1`.
// * **kling-3.0**: Supports `9:16`, `16:9`, `1:1`.
// * **sora-2**: Supports `9:16`, `16:9`.
// * **veo3.1**: Supports `9:16`, `16:9`.
// * **kling-1.6**: Supports `9:16`, `16:9`, `1:1`.
type V1TextToVideoCreateBodyAspectRatioEnum string

const (
	V1TextToVideoCreateBodyAspectRatioEnum169 V1TextToVideoCreateBodyAspectRatioEnum = "16:9"
	V1TextToVideoCreateBodyAspectRatioEnum11  V1TextToVideoCreateBodyAspectRatioEnum = "1:1"
	V1TextToVideoCreateBodyAspectRatioEnum916 V1TextToVideoCreateBodyAspectRatioEnum = "9:16"
)
