package types

import (
	nullable "github.com/magichourhq/magic-hour-go/nullable"
)

// V1SavedItemsListResponseItemsItem
type V1SavedItemsListResponseItemsItem struct {
	Assets []V1SavedItemsListResponseItemsItemAssetsItem `json:"assets"`
	// Unique ID of the saved item.
	Id string `json:"id"`
	// User-provided name of the saved item.
	Name nullable.Nullable[string] `json:"name,omitempty"`
	// Saved item type.
	Type V1SavedItemsListResponseItemsItemTypeEnum `json:"type"`
}
