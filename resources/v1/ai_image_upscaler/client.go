package ai_image_upscaler

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

// AI Image Upscaler
//
// Upscale your image using AI. Each 2x upscale costs 50 frames, and 4x upscale costs 200 frames.
//
// POST /v1/ai-image-upscaler
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiImageUpscalerCreateResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"ai-image-upscaler")
	if err != nil {
		return types.V1AiImageUpscalerCreateResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.V1AiImageUpscalerCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AiImageUpscalerCreateBody{
		Name:        request.Name,
		Assets:      request.Assets,
		ScaleFactor: request.ScaleFactor,
		Style:       request.Style,
	})
	if err != nil {
		return types.V1AiImageUpscalerCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiImageUpscalerCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiImageUpscalerCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiImageUpscalerCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiImageUpscalerCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiImageUpscalerCreateResponse{}, err
	}
	var bodyData types.V1AiImageUpscalerCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiImageUpscalerCreateResponse{}, err
	}
	return bodyData, nil

}
