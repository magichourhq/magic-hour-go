
package face_swap_photo
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

// Face Swap Photo
// 
// Create a face swap photo. Each photo costs 5 frames. The height/width of the output image depends on your subscription. Please refer to our [pricing](/pricing) page for more details
// 
// POST /v1/face-swap-photo
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.PostV1FaceSwapPhotoResponse, error) {
// URL formatting
joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/" + "face-swap-photo")
if err != nil {
return types.PostV1FaceSwapPhotoResponse{}, err
}
targetUrl, err := url.Parse(joinedUrl)
if err != nil {
return types.PostV1FaceSwapPhotoResponse{}, err
}

// Prep body
reqBody, err := json.Marshal(types.PostV1FaceSwapPhotoBody { 
Name: request.Name,
Assets: request.Assets, })
if err != nil {
return types.PostV1FaceSwapPhotoResponse{}, err
}
reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

// Init request
req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
if err != nil {
return types.PostV1FaceSwapPhotoResponse{}, err
}

// Add headers
req.Header.Add("x-sideko-sdk-language", "Go")
req.Header.Add("Content-Type", "application/json")

// Add auth
c.coreClient.AddAuth([]string{"bearerAuth"}, req)

// Add base client & request level modifiers
if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
return types.PostV1FaceSwapPhotoResponse{}, err
}

// Dispatch request
resp, err := c.coreClient.HttpClient.Do(req)
if err != nil {
return types.PostV1FaceSwapPhotoResponse{}, err
}

// Check status
if resp.StatusCode >= 300 {
return types.PostV1FaceSwapPhotoResponse{}, sdkcore.NewApiError(*req, *resp)
}

// Handle response
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
if err != nil {
return types.PostV1FaceSwapPhotoResponse{}, err
}
var bodyData types.PostV1FaceSwapPhotoResponse
err = json.Unmarshal(body, &bodyData)
if err != nil {
return types.PostV1FaceSwapPhotoResponse{}, err
}
return bodyData, nil

}
