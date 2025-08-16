package types

// Success
type V1FilesUploadUrlsCreateResponse struct {
	// The list of upload URLs and file paths for the assets. The response array will match the order of items in the request body. Refer to the [Input Files Guide](/integration/input-files) for more details.
	Items []V1FilesUploadUrlsCreateResponseItemsItem `json:"items"`
}
