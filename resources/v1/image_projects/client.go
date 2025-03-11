
package image_projects
import (
sdkcore "github.com/magichourhq/magic-hour-go/core"
http "net/http"
url "net/url"
types "github.com/magichourhq/magic-hour-go/types"
io "io"
json "encoding/json"
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

// Delete image
// 
// Permanently delete the rendered image. This action is not reversible, please be sure before deleting.
// 
// DELETE /v1/image-projects/{id}
func (c *Client) Delete(request DeleteRequest, reqModifiers ...RequestModifier) error {
// URL formatting
joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/" + "image-projects/" + sdkcore.FmtStringParam(request.Id))
if err != nil {
return err
}
targetUrl, err := url.Parse(joinedUrl)
if err != nil {
return err
}

// Init request
req, err := http.NewRequest("DELETE", targetUrl.String(), nil)
if err != nil {
return err
}

// Add headers
req.Header.Add("x-sideko-sdk-language", "Go")

// Add auth
c.coreClient.AddAuth([]string{"bearerAuth"}, req)

// Add base client & request level modifiers
if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
return err
}

// Dispatch request
resp, err := c.coreClient.HttpClient.Do(req)
if err != nil {
return err
}

// Check status
if resp.StatusCode >= 300 {
return sdkcore.NewApiError(*req, *resp)
}

// No expected response data
return nil

}

// Get image details
// 
// Get the details of a image project. The `download` field will be `null` unless the image was successfully rendered.
// 
// The image can be one of the following status
// - `draft` - not currently used
// - `queued` - the job is queued and waiting for a GPU
// - `rendering` - the generation is in progress
// - `complete` - the image is successful created
// - `error` - an error occurred during rendering
// - `canceled` - image render is canceled by the user
// 
// 
// GET /v1/image-projects/{id}
func (c *Client) Get(request GetRequest, reqModifiers ...RequestModifier) (types.GetV1ImageProjectsIdResponse, error) {
// URL formatting
joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/" + "image-projects/" + sdkcore.FmtStringParam(request.Id))
if err != nil {
return types.GetV1ImageProjectsIdResponse{}, err
}
targetUrl, err := url.Parse(joinedUrl)
if err != nil {
return types.GetV1ImageProjectsIdResponse{}, err
}

// Init request
req, err := http.NewRequest("GET", targetUrl.String(), nil)
if err != nil {
return types.GetV1ImageProjectsIdResponse{}, err
}

// Add headers
req.Header.Add("x-sideko-sdk-language", "Go")

// Add auth
c.coreClient.AddAuth([]string{"bearerAuth"}, req)

// Add base client & request level modifiers
if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
return types.GetV1ImageProjectsIdResponse{}, err
}

// Dispatch request
resp, err := c.coreClient.HttpClient.Do(req)
if err != nil {
return types.GetV1ImageProjectsIdResponse{}, err
}

// Check status
if resp.StatusCode >= 300 {
return types.GetV1ImageProjectsIdResponse{}, sdkcore.NewApiError(*req, *resp)
}

// Handle response
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
if err != nil {
return types.GetV1ImageProjectsIdResponse{}, err
}
var bodyData types.GetV1ImageProjectsIdResponse
err = json.Unmarshal(body, &bodyData)
if err != nil {
return types.GetV1ImageProjectsIdResponse{}, err
}
return bodyData, nil

}
