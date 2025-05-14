package animation

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

// Animation
//
// Create a Animation video. The estimated frame cost is calculated based on the `fps` and `end_seconds` input.
//
// POST /v1/animation
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AnimationCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "animation")
	if err != nil {
		return types.V1AnimationCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AnimationCreateBody{
		Name:       request.Name,
		Assets:     request.Assets,
		EndSeconds: request.EndSeconds,
		Fps:        request.Fps,
		Height:     request.Height,
		Style:      request.Style,
		Width:      request.Width,
	})
	if err != nil {
		return types.V1AnimationCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AnimationCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AnimationCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AnimationCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AnimationCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AnimationCreateResponse{}, err
	}
	var bodyData types.V1AnimationCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AnimationCreateResponse{}, err
	}
	return bodyData, nil

}
