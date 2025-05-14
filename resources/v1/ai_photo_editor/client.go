package ai_photo_editor

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

// AI Photo Editor
//
// > **NOTE**: this API is still in early development stages, and should be avoided. Please reach out to us if you're interested in this API.
//
// Edit photo using AI. Each photo costs 10 credits.
//
// POST /v1/ai-photo-editor
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiPhotoEditorCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "ai-photo-editor")
	if err != nil {
		return types.V1AiPhotoEditorCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AiPhotoEditorCreateBody{
		Name:       request.Name,
		Steps:      request.Steps,
		Assets:     request.Assets,
		Resolution: request.Resolution,
		Style:      request.Style,
	})
	if err != nil {
		return types.V1AiPhotoEditorCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiPhotoEditorCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiPhotoEditorCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiPhotoEditorCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiPhotoEditorCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiPhotoEditorCreateResponse{}, err
	}
	var bodyData types.V1AiPhotoEditorCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiPhotoEditorCreateResponse{}, err
	}
	return bodyData, nil

}
