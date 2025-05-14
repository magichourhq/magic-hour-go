package ai_meme_generator

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

// AI Meme Generator
//
// Create an AI generated meme. Each meme costs 10 credits.
//
// POST /v1/ai-meme-generator
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiMemeGeneratorCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "ai-meme-generator")
	if err != nil {
		return types.V1AiMemeGeneratorCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AiMemeGeneratorCreateBody{
		Name:  request.Name,
		Style: request.Style,
	})
	if err != nil {
		return types.V1AiMemeGeneratorCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiMemeGeneratorCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiMemeGeneratorCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiMemeGeneratorCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiMemeGeneratorCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiMemeGeneratorCreateResponse{}, err
	}
	var bodyData types.V1AiMemeGeneratorCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiMemeGeneratorCreateResponse{}, err
	}
	return bodyData, nil

}
