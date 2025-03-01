package ai_photo_editor

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

// AI Photo Editor
//
// > **NOTE**: this API is still in early development stages, and should be avoided. Please reach out to us if you're interested in this API.
//
// Edit photo using AI. Each photo costs 10 frames.
//
// POST /v1/ai-photo-editor
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.PostV1AiPhotoEditorResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"ai-photo-editor")
	if err != nil {
		return types.PostV1AiPhotoEditorResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.PostV1AiPhotoEditorResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(types.PostV1AiPhotoEditorBody{
		Name:       request.Name,
		Steps:      request.Steps,
		Assets:     request.Assets,
		Resolution: request.Resolution,
		Style:      request.Style})
	if err != nil {
		return types.PostV1AiPhotoEditorResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.PostV1AiPhotoEditorResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.PostV1AiPhotoEditorResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.PostV1AiPhotoEditorResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.PostV1AiPhotoEditorResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.PostV1AiPhotoEditorResponse{}, err
	}
	var bodyData types.PostV1AiPhotoEditorResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.PostV1AiPhotoEditorResponse{}, err
	}
	return bodyData, nil

}
