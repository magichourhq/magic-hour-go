package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1SavedItemsListResponse
type V1SavedItemsListResponse struct {
	Items []V1SavedItemsListResponseItemsItem `json:"items"`
	// Cursor for the next page, or null when there are no more saved items.
	NextCursor nullable.Nullable[string] `json:"next_cursor,omitempty"`
}
