package types

// * `v1` - more detail, closer prompt adherence, and frame-by-frame previews.
// * `v2` - faster, more consistent, and less noisy.
// * `default` - use the default version for the selected art style.
type V1VideoToVideoCreateBodyStyleVersionEnum string

const (
	V1VideoToVideoCreateBodyStyleVersionEnumDefault V1VideoToVideoCreateBodyStyleVersionEnum = "default"
	V1VideoToVideoCreateBodyStyleVersionEnumV1      V1VideoToVideoCreateBodyStyleVersionEnum = "v1"
	V1VideoToVideoCreateBodyStyleVersionEnumV2      V1VideoToVideoCreateBodyStyleVersionEnum = "v2"
)
