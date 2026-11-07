package types

// * `v1` - more detail, closer prompt adherence, and frame-by-frame previews.
// * `v2` - faster, more consistent, and less noisy.
// * `v3` - consistent styling guided by a caller-supplied reference image. Requires assets.initial_image_path and positive width/height.
// * `default` - use the selected art style's default version, including v3.
type V1VideoToVideoCreateBodyStyleVersionEnum string

const (
	V1VideoToVideoCreateBodyStyleVersionEnumDefault V1VideoToVideoCreateBodyStyleVersionEnum = "default"
	V1VideoToVideoCreateBodyStyleVersionEnumV1      V1VideoToVideoCreateBodyStyleVersionEnum = "v1"
	V1VideoToVideoCreateBodyStyleVersionEnumV2      V1VideoToVideoCreateBodyStyleVersionEnum = "v2"
	V1VideoToVideoCreateBodyStyleVersionEnumV3      V1VideoToVideoCreateBodyStyleVersionEnum = "v3"
)
