package audio

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

// Estimate audio credits
//
// Get the credits an audio create request charges. Pass `tool`, then the field below that the tool's create request uses. Does not create a project or spend credits. It prices the fields passed here. It does not check your plan, your balance or the rest of the create request.
//
// GET /v1/estimate/audio
func (c *Client) List(request ListRequest, reqModifiers ...RequestModifier) (types.V1EstimateAudioListResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "estimate/" + "audio")
	if err != nil {
		return types.V1EstimateAudioListResponse{}, err
	}

	// Query params
	params := targetUrl.Query()
	sdkcore.AddQueryParam(params, "tool", request.Tool, "form", true)
	sdkcore.AddQueryParam(params, "character_count", request.CharacterCount, "form", true)
	sdkcore.AddQueryParam(params, "duration", request.Duration, "form", true)
	targetUrl.RawQuery = params.Encode()

	// Init request
	req, err := http.NewRequest("GET", targetUrl.String(), nil)
	if err != nil {
		return types.V1EstimateAudioListResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1EstimateAudioListResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1EstimateAudioListResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1EstimateAudioListResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1EstimateAudioListResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1EstimateAudioListResponse{}, err
	}
	var bodyData types.V1EstimateAudioListResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1EstimateAudioListResponse{}, err
	}
	return bodyData, nil

}
