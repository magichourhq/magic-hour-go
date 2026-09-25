package types

// Output video resolution. Defaults to 480p. 720p and 1080p require a paid plan.
type V1AiVideoTranslatorCreateBodyResolutionEnum string

const (
	V1AiVideoTranslatorCreateBodyResolutionEnum1080p V1AiVideoTranslatorCreateBodyResolutionEnum = "1080p"
	V1AiVideoTranslatorCreateBodyResolutionEnum480p  V1AiVideoTranslatorCreateBodyResolutionEnum = "480p"
	V1AiVideoTranslatorCreateBodyResolutionEnum720p  V1AiVideoTranslatorCreateBodyResolutionEnum = "720p"
)
