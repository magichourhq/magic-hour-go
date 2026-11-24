package checkout

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

// Create checkout
//
// Create a Stripe Checkout link for the account that owns the API key. Nothing is charged until the checkout is completed in a browser.
//
// POST /v1/account/checkout
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AccountCheckoutCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "account/" + "checkout")
	if err != nil {
		return types.V1AccountCheckoutCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AccountCheckoutCreateBody{
		CreditShortfall: request.CreditShortfall,
		PackCount:       request.PackCount,
		Type:            request.Type,
	})
	if err != nil {
		return types.V1AccountCheckoutCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AccountCheckoutCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1AccountCheckoutCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AccountCheckoutCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AccountCheckoutCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AccountCheckoutCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AccountCheckoutCreateResponse{}, err
	}
	var bodyData types.V1AccountCheckoutCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AccountCheckoutCreateResponse{}, err
	}
	return bodyData, nil

}
