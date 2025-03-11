
package types



// * `v1` - more detail, closer prompt adherence, and frame-by-frame previews.
// * `v2` - faster, more consistent, and less noisy.
// * `default` - use the default version for the selected art style.
type PostV1VideoToVideoBodyStyleVersionEnum string
const (
    PostV1VideoToVideoBodyStyleVersionEnumDefault PostV1VideoToVideoBodyStyleVersionEnum = "default"
    PostV1VideoToVideoBodyStyleVersionEnumV1 PostV1VideoToVideoBodyStyleVersionEnum = "v1"
    PostV1VideoToVideoBodyStyleVersionEnumV2 PostV1VideoToVideoBodyStyleVersionEnum = "v2"
)


