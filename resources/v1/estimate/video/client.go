package video

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

// Estimate video credits
//
// Get the credits a video create request charges when it is submitted. Pass `tool`, then the fields below that the tool's create request uses, with the same values. Tools billed per frame charge an estimate at submission, assuming 30 fps for an uploaded video, and settle to the frames that render. Does not create a project or spend credits. It prices the fields passed here. It does not check your plan, your balance or the rest of the create request.
//
// GET /v1/estimate/video
func (c *Client) List(request ListRequest, reqModifiers ...RequestModifier) (types.V1EstimateVideoListResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "estimate/" + "video")
	if err != nil {
		return types.V1EstimateVideoListResponse{}, err
	}

	// Query params
	params := targetUrl.Query()
	sdkcore.AddQueryParam(params, "end_seconds", request.EndSeconds, "form", true)
	sdkcore.AddQueryParam(params, "tool", request.Tool, "form", true)
	sdkcore.AddQueryParam(params, "audio", request.Audio, "form", true)
	sdkcore.AddQueryParam(params, "has_video_reference", request.HasVideoReference, "form", true)
	sdkcore.AddQueryParam(params, "model", request.Model, "form", true)
	sdkcore.AddQueryParam(params, "resolution", request.Resolution, "form", true)
	sdkcore.AddQueryParam(params, "start_seconds", request.StartSeconds, "form", true)
	targetUrl.RawQuery = params.Encode()

	// Init request
	req, err := http.NewRequest("GET", targetUrl.String(), nil)
	if err != nil {
		return types.V1EstimateVideoListResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1EstimateVideoListResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1EstimateVideoListResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1EstimateVideoListResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1EstimateVideoListResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1EstimateVideoListResponse{}, err
	}
	var bodyData types.V1EstimateVideoListResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1EstimateVideoListResponse{}, err
	}
	return bodyData, nil

}
