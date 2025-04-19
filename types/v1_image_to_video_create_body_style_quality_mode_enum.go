package types

// * `quick` - Fastest option for rapid results. Takes ~3 minutes per 5s of video.
// *  `studio` - Polished visuals with longer runtime. Takes ~8.5 minutes per 5s of video.
type V1ImageToVideoCreateBodyStyleQualityModeEnum string

const (
	V1ImageToVideoCreateBodyStyleQualityModeEnumQuick  V1ImageToVideoCreateBodyStyleQualityModeEnum = "quick"
	V1ImageToVideoCreateBodyStyleQualityModeEnumStudio V1ImageToVideoCreateBodyStyleQualityModeEnum = "studio"
)
