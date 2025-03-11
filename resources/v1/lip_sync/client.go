
package lip_sync
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

// Lip Sync
// 
// Create a Lip Sync video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
//   
// Get more information about this mode at our [product page](/products/lip-sync).
//   
// 
// POST /v1/lip-sync
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.PostV1LipSyncResponse, error) {
// URL formatting
joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/" + "lip-sync")
if err != nil {
return types.PostV1LipSyncResponse{}, err
}
targetUrl, err := url.Parse(joinedUrl)
if err != nil {
return types.PostV1LipSyncResponse{}, err
}

// Prep body
reqBody, err := json.Marshal(types.PostV1LipSyncBody { 
MaxFpsLimit: request.MaxFpsLimit,
Name: request.Name,
Assets: request.Assets,
EndSeconds: request.EndSeconds,
Height: request.Height,
StartSeconds: request.StartSeconds,
Width: request.Width, })
if err != nil {
return types.PostV1LipSyncResponse{}, err
}
reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

// Init request
req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
if err != nil {
return types.PostV1LipSyncResponse{}, err
}

// Add headers
req.Header.Add("x-sideko-sdk-language", "Go")
req.Header.Add("Content-Type", "application/json")

// Add auth
c.coreClient.AddAuth([]string{"bearerAuth"}, req)

// Add base client & request level modifiers
if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
return types.PostV1LipSyncResponse{}, err
}

// Dispatch request
resp, err := c.coreClient.HttpClient.Do(req)
if err != nil {
return types.PostV1LipSyncResponse{}, err
}

// Check status
if resp.StatusCode >= 300 {
return types.PostV1LipSyncResponse{}, sdkcore.NewApiError(*req, *resp)
}

// Handle response
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
if err != nil {
return types.PostV1LipSyncResponse{}, err
}
var bodyData types.PostV1LipSyncResponse
err = json.Unmarshal(body, &bodyData)
if err != nil {
return types.PostV1LipSyncResponse{}, err
}
return bodyData, nil

}
