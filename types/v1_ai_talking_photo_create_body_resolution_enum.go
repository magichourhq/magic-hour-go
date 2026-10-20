package types

// Output preset. Overrides `max_resolution`. Defaults to 480p in prompted mode; realistic mode uses the mapped plan maximum.
type V1AiTalkingPhotoCreateBodyResolutionEnum string

const (
	V1AiTalkingPhotoCreateBodyResolutionEnum480p V1AiTalkingPhotoCreateBodyResolutionEnum = "480p"
	V1AiTalkingPhotoCreateBodyResolutionEnum720p V1AiTalkingPhotoCreateBodyResolutionEnum = "720p"
)
