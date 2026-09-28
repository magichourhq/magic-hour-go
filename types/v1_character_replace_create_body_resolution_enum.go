package types

// Output video resolution. Must be supported by `model`. Defaults to the lowest resolution available on your plan for that model.
type V1CharacterReplaceCreateBodyResolutionEnum string

const (
	V1CharacterReplaceCreateBodyResolutionEnum1080p V1CharacterReplaceCreateBodyResolutionEnum = "1080p"
	V1CharacterReplaceCreateBodyResolutionEnum480p  V1CharacterReplaceCreateBodyResolutionEnum = "480p"
	V1CharacterReplaceCreateBodyResolutionEnum720p  V1CharacterReplaceCreateBodyResolutionEnum = "720p"
)
