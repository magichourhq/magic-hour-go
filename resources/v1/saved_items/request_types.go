package saved_items

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
	types "github.com/magichourhq/magic-hour-go/types"
)

// ListRequest
type ListRequest struct {
	// Opaque pagination cursor from the previous response's next_cursor.
	Cursor nullable.Nullable[string] `json:"cursor,omitempty"`
	// Maximum number of saved items to return. Defaults to 20.
	Limit nullable.Nullable[int] `json:"limit,omitempty"`
	// Only return saved items of this type.
	Type nullable.Nullable[types.V1SavedItemsListTypeEnum] `json:"type,omitempty"`
}
