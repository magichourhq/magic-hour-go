package types

// Controls the output video resolution. Defaults to `720p` if not specified.
//
// * **Default**: Supports `480p`, `720p`, and `1080p`.
// * **seedance**: Supports `480p`, `720p`, `1080p`.
// * **kling-2.5**: Supports `720p`, `1080p`.
// * **kling-3.0**: Supports `720p`, `1080p`.
// * **sora-2**: Supports `720p`.
// * **veo3.1**: Supports `720p`, `1080p`.
// * **kling-1.6**: Supports `720p`, `1080p`.
type V1TextToVideoCreateBodyResolutionEnum string

const (
	V1TextToVideoCreateBodyResolutionEnum1080p V1TextToVideoCreateBodyResolutionEnum = "1080p"
	V1TextToVideoCreateBodyResolutionEnum480p  V1TextToVideoCreateBodyResolutionEnum = "480p"
	V1TextToVideoCreateBodyResolutionEnum720p  V1TextToVideoCreateBodyResolutionEnum = "720p"
)
