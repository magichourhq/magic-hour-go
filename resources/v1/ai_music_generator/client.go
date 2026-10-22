package ai_music_generator

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

// AI Music Generator
//
// Generate music from a description, with optional custom lyrics or instrumental-only output. Costs 1 credits per requested second. Returns a project ID; poll GET /v1/audio-projects/{id} for status and downloads.
//
// POST /v1/ai-music-generator
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiMusicGeneratorCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "ai-music-generator")
	if err != nil {
		return types.V1AiMusicGeneratorCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AiMusicGeneratorCreateBody{
		AdditionalProperties: request.AdditionalProperties,
		Name:                 request.Name,
		Duration:             request.Duration,
		Style:                request.Style,
	})
	if err != nil {
		return types.V1AiMusicGeneratorCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiMusicGeneratorCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1AiMusicGeneratorCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiMusicGeneratorCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiMusicGeneratorCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiMusicGeneratorCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiMusicGeneratorCreateResponse{}, err
	}
	var bodyData types.V1AiMusicGeneratorCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiMusicGeneratorCreateResponse{}, err
	}
	return bodyData, nil

}
