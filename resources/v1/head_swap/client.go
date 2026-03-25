package head_swap

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

// Head Swap
//
// Swap a head onto a body image. Each image costs 10 credits. Output resolution depends on your subscription; you may set `max_resolution` lower than your plan maximum if desired.
//
// POST /v1/head-swap
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1HeadSwapCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "head-swap")
	if err != nil {
		return types.V1HeadSwapCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1HeadSwapCreateBody{
		MaxResolution: request.MaxResolution,
		Name:          request.Name,
		Assets:        request.Assets,
	})
	if err != nil {
		return types.V1HeadSwapCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1HeadSwapCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1HeadSwapCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1HeadSwapCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1HeadSwapCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1HeadSwapCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1HeadSwapCreateResponse{}, err
	}
	var bodyData types.V1HeadSwapCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1HeadSwapCreateResponse{}, err
	}
	return bodyData, nil

}
