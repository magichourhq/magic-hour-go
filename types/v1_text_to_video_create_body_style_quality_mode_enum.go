package types

// * `quick` - Fastest option for rapid results. Takes ~3 minutes per 5s of video.
// *  `studio` - Polished visuals with longer runtime. Takes ~8.5 minutes per 5s of video.
type V1TextToVideoCreateBodyStyleQualityModeEnum string

const (
	V1TextToVideoCreateBodyStyleQualityModeEnumQuick  V1TextToVideoCreateBodyStyleQualityModeEnum = "quick"
	V1TextToVideoCreateBodyStyleQualityModeEnumStudio V1TextToVideoCreateBodyStyleQualityModeEnum = "studio"
)
