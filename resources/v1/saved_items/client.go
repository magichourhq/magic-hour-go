package saved_items

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

// List saved items
//
// Returns active saved items owned by the authenticated account, newest first. Each item includes every saved asset with a durable file_path for reuse in compatible generation APIs and a temporary signed URL for previewing or downloading. Filter by type to find characters, references, voices, moodboards, or brand kits. To fetch the next page, pass the response's next_cursor as cursor.
//
// GET /v1/saved-items
func (c *Client) List(request ListRequest, reqModifiers ...RequestModifier) (types.V1SavedItemsListResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "saved-items")
	if err != nil {
		return types.V1SavedItemsListResponse{}, err
	}

	// Query params
	params := targetUrl.Query()
	sdkcore.AddQueryParam(params, "cursor", request.Cursor, "form", true)
	sdkcore.AddQueryParam(params, "limit", request.Limit, "form", true)
	sdkcore.AddQueryParam(params, "type", request.Type, "form", true)
	targetUrl.RawQuery = params.Encode()

	// Init request
	req, err := http.NewRequest("GET", targetUrl.String(), nil)
	if err != nil {
		return types.V1SavedItemsListResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1SavedItemsListResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1SavedItemsListResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1SavedItemsListResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1SavedItemsListResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1SavedItemsListResponse{}, err
	}
	var bodyData types.V1SavedItemsListResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1SavedItemsListResponse{}, err
	}
	return bodyData, nil

}
