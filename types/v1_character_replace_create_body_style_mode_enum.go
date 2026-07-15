package types

// Processing mode. `replace` swaps the detected subject with your reference character. `animate` transfers motion from the video onto your character image.
type V1CharacterReplaceCreateBodyStyleModeEnum string

const (
	V1CharacterReplaceCreateBodyStyleModeEnumAnimate V1CharacterReplaceCreateBodyStyleModeEnum = "animate"
	V1CharacterReplaceCreateBodyStyleModeEnumReplace V1CharacterReplaceCreateBodyStyleModeEnum = "replace"
)
