package ai_image_generator

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

// Create AI Images
//
// Create an AI image. Each image costs 5 frames.
//
// POST /v1/ai-image-generator
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.PostV1AiImageGeneratorResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"ai-image-generator")
	if err != nil {
		return types.PostV1AiImageGeneratorResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.PostV1AiImageGeneratorResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(types.PostV1AiImageGeneratorBody{
		Name:        request.Name,
		ImageCount:  request.ImageCount,
		Orientation: request.Orientation,
		Style:       request.Style})
	if err != nil {
		return types.PostV1AiImageGeneratorResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.PostV1AiImageGeneratorResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.PostV1AiImageGeneratorResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.PostV1AiImageGeneratorResponse{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.PostV1AiImageGeneratorResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.PostV1AiImageGeneratorResponse{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.PostV1AiImageGeneratorResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.PostV1AiImageGeneratorResponse{}, err
	}
	return bodyData, nil

}
