package video_to_video

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

// Video-to-Video
//
// Create a Video To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
//
// Get more information about this mode at our [product page](/products/video-to-video).
//
// POST /v1/video-to-video
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.PostV1VideoToVideoResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"video-to-video")
	if err != nil {
		return types.PostV1VideoToVideoResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.PostV1VideoToVideoResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(types.PostV1VideoToVideoBody{
		FpsResolution: request.FpsResolution,
		Name:          request.Name,
		Assets:        request.Assets,
		EndSeconds:    request.EndSeconds,
		Height:        request.Height,
		StartSeconds:  request.StartSeconds,
		Style:         request.Style,
		Width:         request.Width})
	if err != nil {
		return types.PostV1VideoToVideoResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.PostV1VideoToVideoResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.PostV1VideoToVideoResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.PostV1VideoToVideoResponse{}, err
	}
	defer resp.Body.Close()

	// Handle response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.PostV1VideoToVideoResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.PostV1VideoToVideoResponse{}, sdkcore.NewApiError(*req, *resp, body)
	}
	var bodyData types.PostV1VideoToVideoResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.PostV1VideoToVideoResponse{}, err
	}
	return bodyData, nil

}
