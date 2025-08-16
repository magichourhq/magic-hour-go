package upload_urls

import (
	types "github.com/magichourhq/magic-hour-go/types"
)

// CreateRequest
type CreateRequest struct {
	// The list of assets to upload. The response array will match the order of items in the request body.
	Items []types.V1FilesUploadUrlsCreateBodyItemsItem `json:"items"`
}
