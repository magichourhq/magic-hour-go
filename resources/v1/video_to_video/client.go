package video_to_video

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

// Video-to-Video
//
// **What this API does**
//
// Create the same Video To Video you can make in the browser, but programmatically, so you can automate it, run it at scale, or connect it to your own app or workflow.
//
// **Good for**
// - Automation and batch processing
// - Adding video to video into apps, pipelines, or tools
//
// **How it works (3 steps)**
// 1) Upload your inputs (video, image, or audio) with [Generate Upload URLs](https://docs.magichour.ai/api-reference/files/generate-asset-upload-urls) and copy the `file_path`.
// 2) Send a request to create a video to video job with the basic fields.
// 3) Check the job status until it's `complete`, then download the result from `downloads`.
//
// **Key options**
// - Inputs: usually a file, sometimes a YouTube link, depending on project type
// - Resolution: free users are limited to 576px; higher plans unlock HD and larger sizes
// - Extra fields: e.g. `face_swap_mode`, `start_seconds`/`end_seconds`, or a text prompt
//
// **Cost**
// Credits are only charged for the frames that actually render. You'll see an estimate when the job is queued, and the final total after it's done.
//
// For detailed examples, see the [product page](https://magichour.ai/products/video-to-video).
//
// POST /v1/video-to-video
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1VideoToVideoCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "video-to-video")
	if err != nil {
		return types.V1VideoToVideoCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1VideoToVideoCreateBody{
		FpsResolution: request.FpsResolution,
		Height:        request.Height,
		Name:          request.Name,
		Width:         request.Width,
		Assets:        request.Assets,
		EndSeconds:    request.EndSeconds,
		StartSeconds:  request.StartSeconds,
		Style:         request.Style,
	})
	if err != nil {
		return types.V1VideoToVideoCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1VideoToVideoCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1VideoToVideoCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1VideoToVideoCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1VideoToVideoCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1VideoToVideoCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1VideoToVideoCreateResponse{}, err
	}
	var bodyData types.V1VideoToVideoCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1VideoToVideoCreateResponse{}, err
	}
	return bodyData, nil

}
