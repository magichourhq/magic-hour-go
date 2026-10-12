package types

// * `v1` - more detail, closer prompt adherence, and frame-by-frame previews.
// * `v2` - faster, more consistent, and less noisy.
// * `default` - use the selected art style's default version. Some art styles default to v3, which is unavailable through this API. For `Studio Ghibli`, explicitly set `version` to `v2` to use its supported preset.
type V1VideoToVideoCreateBodyStyleVersionEnum string

const (
	V1VideoToVideoCreateBodyStyleVersionEnumDefault V1VideoToVideoCreateBodyStyleVersionEnum = "default"
	V1VideoToVideoCreateBodyStyleVersionEnumV1      V1VideoToVideoCreateBodyStyleVersionEnum = "v1"
	V1VideoToVideoCreateBodyStyleVersionEnumV2      V1VideoToVideoCreateBodyStyleVersionEnum = "v2"
)
