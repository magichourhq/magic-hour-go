
package files
import (
sdkcore "github.com/magichourhq/magic-hour-go/core"
http "net/http"
upload_urls "github.com/magichourhq/magic-hour-go/resources/v1/files/upload_urls"
)
type Client struct {
	coreClient *sdkcore.CoreClient
UploadUrls *upload_urls.Client
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient: coreClient,
UploadUrls: upload_urls.NewClient(coreClient),
	}

	return &client
}
