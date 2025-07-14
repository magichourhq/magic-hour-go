package types

// DEPRECATED: Please use `resolution` field instead. For backward compatibility:
// * `quick` maps to 720p resolution
// * `studio` maps to 1080p resolution
//
// This field will be removed in a future version. Use the `resolution` field to directly to specify the resolution.
type V1ImageToVideoCreateBodyStyleQualityModeEnum string

const (
	V1ImageToVideoCreateBodyStyleQualityModeEnumQuick  V1ImageToVideoCreateBodyStyleQualityModeEnum = "quick"
	V1ImageToVideoCreateBodyStyleQualityModeEnumStudio V1ImageToVideoCreateBodyStyleQualityModeEnum = "studio"
)
