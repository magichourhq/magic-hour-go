package face_swap_photo

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

// Face Swap Photo
//
// Create a face swap photo. Each photo costs 5 credits. The height/width of the output image depends on your subscription. Please refer to our [pricing](https://magichour.ai/pricing) page for more details
//
// POST /v1/face-swap-photo
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1FaceSwapPhotoCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "face-swap-photo")
	if err != nil {
		return types.V1FaceSwapPhotoCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1FaceSwapPhotoCreateBody{
		Name:   request.Name,
		Assets: request.Assets,
	})
	if err != nil {
		return types.V1FaceSwapPhotoCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1FaceSwapPhotoCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1FaceSwapPhotoCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1FaceSwapPhotoCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1FaceSwapPhotoCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1FaceSwapPhotoCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1FaceSwapPhotoCreateResponse{}, err
	}
	var bodyData types.V1FaceSwapPhotoCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1FaceSwapPhotoCreateResponse{}, err
	}
	return bodyData, nil

}
