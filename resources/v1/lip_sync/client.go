package lip_sync

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

// Lip Sync
//
// Create a Lip Sync video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
//
// Get more information about this mode at our [product page](https://magichour.ai/products/lip-sync).
//
// POST /v1/lip-sync
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1LipSyncCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "lip-sync")
	if err != nil {
		return types.V1LipSyncCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1LipSyncCreateBody{
		Height:       request.Height,
		MaxFpsLimit:  request.MaxFpsLimit,
		Name:         request.Name,
		Style:        request.Style,
		Width:        request.Width,
		Assets:       request.Assets,
		EndSeconds:   request.EndSeconds,
		StartSeconds: request.StartSeconds,
	})
	if err != nil {
		return types.V1LipSyncCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1LipSyncCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1LipSyncCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1LipSyncCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1LipSyncCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1LipSyncCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1LipSyncCreateResponse{}, err
	}
	var bodyData types.V1LipSyncCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1LipSyncCreateResponse{}, err
	}
	return bodyData, nil

}
