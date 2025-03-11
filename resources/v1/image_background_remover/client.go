
package image_background_remover
import (
sdkcore "github.com/magichourhq/magic-hour-go/core"
http "net/http"
types "github.com/magichourhq/magic-hour-go/types"
url "net/url"
json "encoding/json"
bytes "bytes"
io "io"
)
type Client struct {
	coreClient *sdkcore.CoreClient
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient: coreClient,
	}

	return &client
}

// Image Background Remover
// 
// Remove background from image. Each image costs 5 frames.
// 
// POST /v1/image-background-remover
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.PostV1ImageBackgroundRemoverResponse, error) {
// URL formatting
joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/" + "image-background-remover")
if err != nil {
return types.PostV1ImageBackgroundRemoverResponse{}, err
}
targetUrl, err := url.Parse(joinedUrl)
if err != nil {
return types.PostV1ImageBackgroundRemoverResponse{}, err
}

// Prep body
reqBody, err := json.Marshal(types.PostV1ImageBackgroundRemoverBody { 
Name: request.Name,
Assets: request.Assets, })
if err != nil {
return types.PostV1ImageBackgroundRemoverResponse{}, err
}
reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

// Init request
req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
if err != nil {
return types.PostV1ImageBackgroundRemoverResponse{}, err
}

// Add headers
req.Header.Add("x-sideko-sdk-language", "Go")
req.Header.Add("Content-Type", "application/json")

// Add auth
c.coreClient.AddAuth([]string{"bearerAuth"}, req)

// Add base client & request level modifiers
if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
return types.PostV1ImageBackgroundRemoverResponse{}, err
}

// Dispatch request
resp, err := c.coreClient.HttpClient.Do(req)
if err != nil {
return types.PostV1ImageBackgroundRemoverResponse{}, err
}

// Check status
if resp.StatusCode >= 300 {
return types.PostV1ImageBackgroundRemoverResponse{}, sdkcore.NewApiError(*req, *resp)
}

// Handle response
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
if err != nil {
return types.PostV1ImageBackgroundRemoverResponse{}, err
}
var bodyData types.PostV1ImageBackgroundRemoverResponse
err = json.Unmarshal(body, &bodyData)
if err != nil {
return types.PostV1ImageBackgroundRemoverResponse{}, err
}
return bodyData, nil

}
