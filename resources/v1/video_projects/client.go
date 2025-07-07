package video_projects

import (
	json "encoding/json"
	io "io"
	http "net/http"

	sdkcore "github.com/magichourhq/magic-hour-go/core"
	types "github.com/magichourhq/magic-hour-go/types"
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

// Delete video
//
// Permanently delete the rendered video. This action is not reversible, please be sure before deleting.
//
// DELETE /v1/video-projects/{id}
func (c *Client) Delete(request DeleteRequest, reqModifiers ...RequestModifier) error {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "video-projects/" + sdkcore.FmtStringParam(request.Id))
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
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return err
	}

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

// Get video details
//
// Get the details of a video project. The `downloads` field will be empty unless the video was successfully rendered.
//
// The video can be one of the following status
// - `draft` - not currently used
// - `queued` - the job is queued and waiting for a GPU
// - `rendering` - the generation is in progress
// - `complete` - the video is successful created
// - `error` - an error occurred during rendering
// - `canceled` - video render is canceled by the user
//
// GET /v1/video-projects/{id}
func (c *Client) Get(request GetRequest, reqModifiers ...RequestModifier) (types.V1VideoProjectsGetResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "video-projects/" + sdkcore.FmtStringParam(request.Id))
	if err != nil {
		return types.V1VideoProjectsGetResponse{}, err
	}

	// Init request
	req, err := http.NewRequest("GET", targetUrl.String(), nil)
	if err != nil {
		return types.V1VideoProjectsGetResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1VideoProjectsGetResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1VideoProjectsGetResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1VideoProjectsGetResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1VideoProjectsGetResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1VideoProjectsGetResponse{}, err
	}
	var bodyData types.V1VideoProjectsGetResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1VideoProjectsGetResponse{}, err
	}
	return bodyData, nil

}
