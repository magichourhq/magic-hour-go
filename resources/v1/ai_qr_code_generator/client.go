package ai_qr_code_generator

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

// AI QR Code
//
// Create an AI QR code. Each QR code costs 20 frames.
//
// POST /v1/ai-qr-code-generator
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiQrCodeGeneratorcreateResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"ai-qr-code-generator")
	if err != nil {
		return types.V1AiQrCodeGeneratorcreateResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.V1AiQrCodeGeneratorcreateResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(types.PostV1AiQrCodeGeneratorBody{
		Name:    request.Name,
		Content: request.Content,
		Style:   request.Style})
	if err != nil {
		return types.V1AiQrCodeGeneratorcreateResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiQrCodeGeneratorcreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiQrCodeGeneratorcreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiQrCodeGeneratorcreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiQrCodeGeneratorcreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiQrCodeGeneratorcreateResponse{}, err
	}
	var bodyData types.V1AiQrCodeGeneratorcreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiQrCodeGeneratorcreateResponse{}, err
	}
	return bodyData, nil

}
