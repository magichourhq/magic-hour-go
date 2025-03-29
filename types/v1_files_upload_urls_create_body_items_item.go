package types

// V1FilesUploadUrlsCreateBodyItemsItem
type V1FilesUploadUrlsCreateBodyItemsItem struct {
	// the extension of the file to upload. Do not include the dot (.) before the extension.
	Extension string `json:"extension"`
	// The type of asset to upload
	Type V1FilesUploadUrlsCreateBodyItemsItemTypeEnum `json:"type"`
}
