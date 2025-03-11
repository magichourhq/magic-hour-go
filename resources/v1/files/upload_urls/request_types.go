
package upload_urls
import (
types "github.com/magichourhq/magic-hour-go/types"
)
// CreateRequest
type CreateRequest struct {
    Items []types.PostV1FilesUploadUrlsBodyItemsItem `json:"items"`
}



