package text_to_video

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

// Text-to-Video
//
// Create a Text To Video video. The estimated frame cost is calculated using 30 FPS. This amount is deducted from your account balance when a video is queued. Once the video is complete, the cost will be updated based on the actual number of frames rendered.
//
// Get more information about this mode at our [product page](https://magichour.ai/products/text-to-video).
//
// POST /v1/text-to-video
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1TextToVideoCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "text-to-video")
	if err != nil {
		return types.V1TextToVideoCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1TextToVideoCreateBody{
		Name:        request.Name,
		Resolution:  request.Resolution,
		EndSeconds:  request.EndSeconds,
		Orientation: request.Orientation,
		Style:       request.Style,
	})
	if err != nil {
		return types.V1TextToVideoCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1TextToVideoCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1TextToVideoCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1TextToVideoCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1TextToVideoCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1TextToVideoCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1TextToVideoCreateResponse{}, err
	}
	var bodyData types.V1TextToVideoCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1TextToVideoCreateResponse{}, err
	}
	return bodyData, nil

}
