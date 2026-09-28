package types

// Model to use. Defaults to `wan-animate`.
//
// * **`wan-animate`**: 480p, 720p. Supports `points` subject selection.
// * **`kling-3.0`**: 720p, 1080p. Clips of 3–10 seconds in `replace` mode or 3–30 seconds in `animate` mode. Picks the main person automatically, so `points` are rejected.
type V1CharacterReplaceCreateBodyModelEnum string

const (
	V1CharacterReplaceCreateBodyModelEnumKling30    V1CharacterReplaceCreateBodyModelEnum = "kling-3.0"
	V1CharacterReplaceCreateBodyModelEnumWanAnimate V1CharacterReplaceCreateBodyModelEnum = "wan-animate"
)
