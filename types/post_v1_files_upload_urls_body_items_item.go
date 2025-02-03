package types

// PostV1FilesUploadUrlsBodyItemsItem
type PostV1FilesUploadUrlsBodyItemsItem struct {
	// the extension of the file to upload. Do not include the dot (.) before the extension.
	Extension string `json:"extension"`
	// The type of asset to upload
	Type PostV1FilesUploadUrlsBodyItemsItemTypeEnum `json:"type"`
}
