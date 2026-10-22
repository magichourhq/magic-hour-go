package types

// Output resolution. `720p` and `1080p` set the output's short side, `2k` and `4k` its long side; the aspect ratio is kept. Requires a paid plan.
type V1AiVideoUpscalerCreateBodyResolutionEnum string

const (
	V1AiVideoUpscalerCreateBodyResolutionEnum1080p V1AiVideoUpscalerCreateBodyResolutionEnum = "1080p"
	V1AiVideoUpscalerCreateBodyResolutionEnum2k    V1AiVideoUpscalerCreateBodyResolutionEnum = "2k"
	V1AiVideoUpscalerCreateBodyResolutionEnum4k    V1AiVideoUpscalerCreateBodyResolutionEnum = "4k"
	V1AiVideoUpscalerCreateBodyResolutionEnum720p  V1AiVideoUpscalerCreateBodyResolutionEnum = "720p"
)
