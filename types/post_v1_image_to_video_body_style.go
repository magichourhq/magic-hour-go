package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// PostV1ImageToVideoBodyStyle
type PostV1ImageToVideoBodyStyle struct {
	// The prompt used for the video.
	Prompt nullable.Nullable[string] `json:"prompt,omitempty"`
}
