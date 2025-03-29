package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1ImageToVideoCreateBodyStyle
type V1ImageToVideoCreateBodyStyle struct {
	// High Quality mode enhances detail, sharpness, and realism, making it ideal for portraits, animals, and intricate landscapes.
	HighQuality nullable.Nullable[bool] `json:"high_quality,omitempty"`
	// The prompt used for the video.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
