package types

// Determines whether the resulting video will have the same frame per second as the original video, or half.
// * `FULL` - the result video will have the same FPS as the input video
// * `HALF` - the result video will have half the FPS as the input video
type V1VideoToVideoCreateBodyFpsResolutionEnum string

const (
	V1VideoToVideoCreateBodyFpsResolutionEnumFull V1VideoToVideoCreateBodyFpsResolutionEnum = "FULL"
	V1VideoToVideoCreateBodyFpsResolutionEnumHalf V1VideoToVideoCreateBodyFpsResolutionEnum = "HALF"
)
