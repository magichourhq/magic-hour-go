package types

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
type V1ImageToVideoCreateBodyResolutionEnum string

const (
	V1ImageToVideoCreateBodyResolutionEnum1080p V1ImageToVideoCreateBodyResolutionEnum = "1080p"
	V1ImageToVideoCreateBodyResolutionEnum480p  V1ImageToVideoCreateBodyResolutionEnum = "480p"
	V1ImageToVideoCreateBodyResolutionEnum4k    V1ImageToVideoCreateBodyResolutionEnum = "4k"
	V1ImageToVideoCreateBodyResolutionEnum720p  V1ImageToVideoCreateBodyResolutionEnum = "720p"
)
