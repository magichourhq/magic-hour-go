package ai_voice_generator

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

// AI Voice Generator
//
// Generate speech from text. Each character costs 0.05 credits. The cost is rounded up to the nearest whole number.
//
// POST /v1/ai-voice-generator
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiVoiceGeneratorCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "ai-voice-generator")
	if err != nil {
		return types.V1AiVoiceGeneratorCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AiVoiceGeneratorCreateBody{
		Name:  request.Name,
		Style: request.Style,
	})
	if err != nil {
		return types.V1AiVoiceGeneratorCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiVoiceGeneratorCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1AiVoiceGeneratorCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiVoiceGeneratorCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiVoiceGeneratorCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiVoiceGeneratorCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiVoiceGeneratorCreateResponse{}, err
	}
	var bodyData types.V1AiVoiceGeneratorCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiVoiceGeneratorCreateResponse{}, err
	}
	return bodyData, nil

}
