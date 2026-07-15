package character_replace

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	Data nullable.Nullable[types.V1CharacterReplaceCreateBody] `json:"data,omitempty"`
}
