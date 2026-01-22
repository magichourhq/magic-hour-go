package types

// Determines the aspect ratio of the output video.
// * **Seedance**: Supports `9:16`, `16:9`, `1:1`.
// * **Kling 2.5 Audio**: Supports `9:16`, `16:9`, `1:1`.
// * **Sora 2**: Supports `9:16`, `16:9`.
// * **Veo 3.1 Audio**: Supports `9:16`, `16:9`.
// * **Veo 3.1**: Supports `9:16`, `16:9`.
// * **Kling 1.6**: Supports `9:16`, `16:9`, `1:1`.
type V1TextToVideoCreateBodyAspectRatioEnum string

const (
	V1TextToVideoCreateBodyAspectRatioEnum169 V1TextToVideoCreateBodyAspectRatioEnum = "16:9"
	V1TextToVideoCreateBodyAspectRatioEnum11  V1TextToVideoCreateBodyAspectRatioEnum = "1:1"
	V1TextToVideoCreateBodyAspectRatioEnum916 V1TextToVideoCreateBodyAspectRatioEnum = "9:16"
)
