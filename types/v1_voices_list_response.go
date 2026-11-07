package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1VoicesListResponse
type V1VoicesListResponse struct {
	// Cursor for the next page, or null when there are no more voices.
	NextCursor nullable.Nullable[string]        `json:"next_cursor,omitempty"`
	Voices     []V1VoicesListResponseVoicesItem `json:"voices"`
}
