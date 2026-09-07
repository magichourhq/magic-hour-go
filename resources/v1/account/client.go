package account

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

// Get account details
//
// Get the current credit balance and subscription details of the account that owns the API key.
//
// GET /v1/account
func (c *Client) List(reqModifiers ...RequestModifier) (types.V1AccountListResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "account")
	if err != nil {
		return types.V1AccountListResponse{}, err
	}

	// Init request
	req, err := http.NewRequest("GET", targetUrl.String(), nil)
	if err != nil {
		return types.V1AccountListResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1AccountListResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AccountListResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AccountListResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AccountListResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AccountListResponse{}, err
	}
	var bodyData types.V1AccountListResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AccountListResponse{}, err
	}
	return bodyData, nil

}
