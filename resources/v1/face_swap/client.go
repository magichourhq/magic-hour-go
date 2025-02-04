package face_swap

import (
	bytes "bytes"
	json "encoding/json"
	io "io"
	http "net/http"
	url "net/url"

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

// Face Swap video
//
// Create a Face Swap video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
//
// Get more information about this mode at our [product page](/products/face-swap).
//
// POST /v1/face-swap
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1FaceSwapcreateResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"face-swap")
	if err != nil {
		return types.V1FaceSwapcreateResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.V1FaceSwapcreateResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(types.PostV1FaceSwapBody{
		Name:         request.Name,
		Assets:       request.Assets,
		EndSeconds:   request.EndSeconds,
		Height:       request.Height,
		StartSeconds: request.StartSeconds,
		Width:        request.Width})
	if err != nil {
		return types.V1FaceSwapcreateResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1FaceSwapcreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1FaceSwapcreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1FaceSwapcreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1FaceSwapcreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1FaceSwapcreateResponse{}, err
	}
	var bodyData types.V1FaceSwapcreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1FaceSwapcreateResponse{}, err
	}
	return bodyData, nil

}
