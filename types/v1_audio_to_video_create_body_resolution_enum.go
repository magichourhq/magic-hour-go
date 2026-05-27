package types

// Output video resolution. Defaults to `720p` on paid tiers and `480p` on free tiers.
type V1AudioToVideoCreateBodyResolutionEnum string

const (
	V1AudioToVideoCreateBodyResolutionEnum1080p V1AudioToVideoCreateBodyResolutionEnum = "1080p"
	V1AudioToVideoCreateBodyResolutionEnum480p  V1AudioToVideoCreateBodyResolutionEnum = "480p"
	V1AudioToVideoCreateBodyResolutionEnum720p  V1AudioToVideoCreateBodyResolutionEnum = "720p"
)
