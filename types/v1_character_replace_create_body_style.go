package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// Optional style controls for replace vs animate mode and subject selection.
type V1CharacterReplaceCreateBodyStyle struct {
	// Processing mode. `replace` swaps the detected subject with your reference character. `animate` transfers motion from the video onto your character image.
	Mode nullable.Nullable[V1CharacterReplaceCreateBodyStyleModeEnum] `json:"mode,omitempty"`
	// On-frame markers for manual subject selection. Required when `selection_mode` is `point`. Ignored when `selection_mode` is `auto` or omitted.
	Points nullable.Nullable[[]V1CharacterReplaceCreateBodyStylePointsItem] `json:"points,omitempty"`
	// How to locate the subject in the source video. `auto` detects a person automatically. `point` uses your `points` to mark the subject. Defaults to `auto`.
	SelectionMode nullable.Nullable[V1CharacterReplaceCreateBodyStyleSelectionModeEnum] `json:"selection_mode,omitempty"`
}
