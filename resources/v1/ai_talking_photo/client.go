package ai_talking_photo

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

// AI Talking Photo
//
// Create a talking photo from an image and audio or text input.
//
// POST /v1/ai-talking-photo
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiTalkingPhotoCreateResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"ai-talking-photo")
	if err != nil {
		return types.V1AiTalkingPhotoCreateResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.V1AiTalkingPhotoCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AiTalkingPhotoCreateBody{
		Name:         request.Name,
		Assets:       request.Assets,
		EndSeconds:   request.EndSeconds,
		StartSeconds: request.StartSeconds,
	})
	if err != nil {
		return types.V1AiTalkingPhotoCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiTalkingPhotoCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiTalkingPhotoCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiTalkingPhotoCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiTalkingPhotoCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiTalkingPhotoCreateResponse{}, err
	}
	var bodyData types.V1AiTalkingPhotoCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiTalkingPhotoCreateResponse{}, err
	}
	return bodyData, nil

}
