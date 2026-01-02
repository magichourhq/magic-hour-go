package ai_voice_cloner

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

// AI Voice Cloner
//
// Clone a voice from an audio sample and generate speech.
// * Each character costs 0.05 credits.
// * The cost is rounded up to the nearest whole number
//
// POST /v1/ai-voice-cloner
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiVoiceClonerCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "ai-voice-cloner")
	if err != nil {
		return types.V1AiVoiceClonerCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AiVoiceClonerCreateBody{
		Name:   request.Name,
		Assets: request.Assets,
		Style:  request.Style,
	})
	if err != nil {
		return types.V1AiVoiceClonerCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiVoiceClonerCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1AiVoiceClonerCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiVoiceClonerCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiVoiceClonerCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiVoiceClonerCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiVoiceClonerCreateResponse{}, err
	}
	var bodyData types.V1AiVoiceClonerCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiVoiceClonerCreateResponse{}, err
	}
	return bodyData, nil

}
