package voices

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

// List voices
//
// List the voices available for text-to-speech. Use a voice's `id` as `style.voice_name` in the AI Voice Generator API. To fetch the next page, pass the response's next_cursor as cursor.
//
// GET /v1/voices
func (c *Client) List(request ListRequest, reqModifiers ...RequestModifier) (types.V1VoicesListResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "voices")
	if err != nil {
		return types.V1VoicesListResponse{}, err
	}

	// Query params
	params := targetUrl.Query()
	sdkcore.AddQueryParam(params, "cursor", request.Cursor, "form", true)
	sdkcore.AddQueryParam(params, "limit", request.Limit, "form", true)
	targetUrl.RawQuery = params.Encode()

	// Init request
	req, err := http.NewRequest("GET", targetUrl.String(), nil)
	if err != nil {
		return types.V1VoicesListResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1VoicesListResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1VoicesListResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1VoicesListResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1VoicesListResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1VoicesListResponse{}, err
	}
	var bodyData types.V1VoicesListResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1VoicesListResponse{}, err
	}
	return bodyData, nil

}
