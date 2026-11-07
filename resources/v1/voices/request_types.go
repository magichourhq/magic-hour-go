package voices

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// ListRequest
type ListRequest struct {
	// Opaque pagination cursor from the previous response's next_cursor.
	Cursor nullable.Nullable[string] `json:"cursor,omitempty"`
	// Maximum number of voices to return. Defaults to 20.
	Limit nullable.Nullable[int] `json:"limit,omitempty"`
}
