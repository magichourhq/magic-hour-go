package types

// Defaults to 720p on paid plans and 480p on free plans. Unsupported model/resolution combinations and resolutions unavailable for your plan are rejected.
type V1AiUgcAdCreateBodyResolutionEnum string

const (
	V1AiUgcAdCreateBodyResolutionEnum1080p V1AiUgcAdCreateBodyResolutionEnum = "1080p"
	V1AiUgcAdCreateBodyResolutionEnum480p  V1AiUgcAdCreateBodyResolutionEnum = "480p"
	V1AiUgcAdCreateBodyResolutionEnum720p  V1AiUgcAdCreateBodyResolutionEnum = "720p"
)
