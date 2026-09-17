package types

// V1SavedItemsListResponseItemsItemAssetsItem
type V1SavedItemsListResponseItemsItemAssetsItem struct {
	// Durable asset path. Pass it to a compatible API asset field without uploading it again.
	FilePath string `json:"file_path"`
	// Whether this asset is the saved item's primary asset.
	IsPrimary bool `json:"is_primary"`
	// Media type of the asset.
	MediaKind V1SavedItemsListResponseItemsItemAssetsItemMediaKindEnum `json:"media_kind"`
	// Signed URL for previewing or downloading the asset. Expires after 24 hours.
	Url string `json:"url"`
	// When the signed URL expires. The saved asset and file_path do not expire.
	UrlExpiresAt string `json:"url_expires_at"`
}
