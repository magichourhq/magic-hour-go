package types

// How to locate the subject in the source video. `auto` detects a person automatically. `point` uses your `points` to mark the subject and is supported by `wan-animate`, `minimax-h3`. `minimax-h3` subject `points` are `replace`-only; `animate` with `selection_mode` `point` returns 422, and with `auto` or no `selection_mode` the points are ignored. Defaults to `auto`.
type V1CharacterReplaceCreateBodyStyleSelectionModeEnum string

const (
	V1CharacterReplaceCreateBodyStyleSelectionModeEnumAuto  V1CharacterReplaceCreateBodyStyleSelectionModeEnum = "auto"
	V1CharacterReplaceCreateBodyStyleSelectionModeEnumPoint V1CharacterReplaceCreateBodyStyleSelectionModeEnum = "point"
)
