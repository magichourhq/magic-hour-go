package types

// Controls the output video resolution. Defaults to `720p` if not specified.
//
// * **Default**: Supports `480p`, `720p`, and `1080p`.
// * **Seedance**: Supports `480p`, `720p`, `1080p`.
// * **Kling 2.5 Audio**: Supports `720p`, `1080p`.
// * **Sora 2**: Supports `720p`.
// * **Veo 3.1 Audio**: Supports `720p`, `1080p`.
// * **Veo 3.1**: Supports `720p`, `1080p`.
// * **Kling 1.6**: Supports `720p`, `1080p`.
type V1TextToVideoCreateBodyResolutionEnum string

const (
	V1TextToVideoCreateBodyResolutionEnum1080p V1TextToVideoCreateBodyResolutionEnum = "1080p"
	V1TextToVideoCreateBodyResolutionEnum480p  V1TextToVideoCreateBodyResolutionEnum = "480p"
	V1TextToVideoCreateBodyResolutionEnum720p  V1TextToVideoCreateBodyResolutionEnum = "720p"
)
