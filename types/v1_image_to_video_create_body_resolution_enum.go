package types

// Controls the output video resolution. Defaults to `720p` if not specified.
//
// **Options:**
// - `480p` - Supports only 5 or 10 second videos. Output: 24fps. Cost: 120 credits per 5 seconds.
// - `720p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 300 credits per 5 seconds.
// - `1080p` - Supports videos between 5-60 seconds. Output: 30fps. Cost: 600 credits per 5 seconds. **Requires** `pro` or `business` tier.
type V1ImageToVideoCreateBodyResolutionEnum string

const (
	V1ImageToVideoCreateBodyResolutionEnum1080p V1ImageToVideoCreateBodyResolutionEnum = "1080p"
	V1ImageToVideoCreateBodyResolutionEnum480p  V1ImageToVideoCreateBodyResolutionEnum = "480p"
	V1ImageToVideoCreateBodyResolutionEnum720p  V1ImageToVideoCreateBodyResolutionEnum = "720p"
)
