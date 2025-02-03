package ai_headshot_generator

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

// AI Headshots
//
// Create an AI headshot. Each headshot costs 50 frames.
//
// POST /v1/ai-headshot-generator
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.PostV1AiHeadshotGeneratorResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"ai-headshot-generator")
	if err != nil {
		return types.PostV1AiHeadshotGeneratorResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.PostV1AiHeadshotGeneratorResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(types.PostV1AiHeadshotGeneratorBody{
		Name:   request.Name,
		Assets: request.Assets})
	if err != nil {
		return types.PostV1AiHeadshotGeneratorResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.PostV1AiHeadshotGeneratorResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.PostV1AiHeadshotGeneratorResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.PostV1AiHeadshotGeneratorResponse{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.PostV1AiHeadshotGeneratorResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.PostV1AiHeadshotGeneratorResponse{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.PostV1AiHeadshotGeneratorResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.PostV1AiHeadshotGeneratorResponse{}, err
	}
	return bodyData, nil

}
