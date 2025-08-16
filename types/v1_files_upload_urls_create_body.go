package types

// V1FilesUploadUrlsCreateBody
type V1FilesUploadUrlsCreateBody struct {
	// The list of assets to upload. The response array will match the order of items in the request body.
	Items []V1FilesUploadUrlsCreateBodyItemsItem `json:"items"`
}
