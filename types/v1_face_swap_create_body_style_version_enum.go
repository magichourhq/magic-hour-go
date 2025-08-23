package types

// * `v1` - May preserve skin detail and texture better, but weaker identity preservation.
// * `v2` - Faster, sharper, better handling of hair and glasses. stronger identity preservation. (Recommended)
// * `default` - Use the version we recommend, which will change over time. This is recommended unless you need a specific earlier version. This is the default behavior.
type V1FaceSwapCreateBodyStyleVersionEnum string

const (
	V1FaceSwapCreateBodyStyleVersionEnumDefault V1FaceSwapCreateBodyStyleVersionEnum = "default"
	V1FaceSwapCreateBodyStyleVersionEnumV1      V1FaceSwapCreateBodyStyleVersionEnum = "v1"
	V1FaceSwapCreateBodyStyleVersionEnumV2      V1FaceSwapCreateBodyStyleVersionEnum = "v2"
)
