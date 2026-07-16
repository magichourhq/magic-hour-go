package types

// How to locate the subject in the source video. `auto` detects a person automatically. `point` uses your `points` to mark the subject. Defaults to `auto`.
type V1CharacterReplaceCreateBodyStyleSelectionModeEnum string

const (
	V1CharacterReplaceCreateBodyStyleSelectionModeEnumAuto  V1CharacterReplaceCreateBodyStyleSelectionModeEnum = "auto"
	V1CharacterReplaceCreateBodyStyleSelectionModeEnumPoint V1CharacterReplaceCreateBodyStyleSelectionModeEnum = "point"
)
