// Code generated by Sideko (sideko.dev) DO NOT EDIT.

package text_to_video

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

// Create a Text To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
//
// Get more information about this mode at our [product page](/products/text-to-video).
//
// POST /v1/text-to-video
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.PostV1TextToVideoResponse, error) {
	// URL formatting
	joined, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"text-to-video")
	if err != nil {
		return types.PostV1TextToVideoResponse{}, err
	}
	url, err := url.Parse(joined)
	if err != nil {
		return types.PostV1TextToVideoResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(request.Data)
	if err != nil {
		return types.PostV1TextToVideoResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", url.String(), reqBodyBuf)
	if err != nil {
		return types.PostV1TextToVideoResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.PostV1TextToVideoResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.PostV1TextToVideoResponse{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.PostV1TextToVideoResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.PostV1TextToVideoResponse{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.PostV1TextToVideoResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.PostV1TextToVideoResponse{}, err
	}
	return bodyData, nil

}
