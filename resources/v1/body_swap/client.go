package body_swap

import (
	bytes "bytes"
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

// Body Swap
//
// Swap a person into a scene image using Nano Banana 2 Lite (640px/1k) or Nano Banana 2 (2k/4k). Credits depend on `resolution` (from 50 credits at 640px upward).
//
// POST /v1/body-swap
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1BodySwapCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "body-swap")
	if err != nil {
		return types.V1BodySwapCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1BodySwapCreateBody{
		Name:       request.Name,
		Assets:     request.Assets,
		Resolution: request.Resolution,
	})
	if err != nil {
		return types.V1BodySwapCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1BodySwapCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1BodySwapCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1BodySwapCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1BodySwapCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1BodySwapCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1BodySwapCreateResponse{}, err
	}
	var bodyData types.V1BodySwapCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1BodySwapCreateResponse{}, err
	}
	return bodyData, nil

}
