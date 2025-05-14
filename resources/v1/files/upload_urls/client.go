package upload_urls

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

// Generate asset upload urls
//
// Create a list of urls used to upload the assets needed to generate a video. Each video type has their own requirements on what assets are required. Please refer to the specific mode API for more details. The response array will be in the same order as the request body.
//
// Below is the list of valid extensions for each asset type:
//
// - video: mp4, m4v, mov, webm
// - audio: mp3, mpeg, wav, aac, aiff, flac
// - image: png, jpg, jpeg, webp, avif, jp2, tiff, bmp
//
// Note: `.gif` is supported for face swap API `video_file_path` field.
//
// After receiving the upload url, you can upload the file by sending a PUT request with the header `'Content-Type: application/octet-stream'`.
//
// # For example using curl
//
// ```
//
//	curl -X PUT -H 'Content-Type: application/octet-stream' \
//	  --data '@/path/to/file/video.mp4' \
//	  https://videos.magichour.ai/api-assets/id/video.mp4?auth-value=1234567890
//
// ```
//
// POST /v1/files/upload-urls
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1FilesUploadUrlsCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "files/" + "upload-urls")
	if err != nil {
		return types.V1FilesUploadUrlsCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1FilesUploadUrlsCreateBody{
		Items: request.Items,
	})
	if err != nil {
		return types.V1FilesUploadUrlsCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1FilesUploadUrlsCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1FilesUploadUrlsCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1FilesUploadUrlsCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1FilesUploadUrlsCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1FilesUploadUrlsCreateResponse{}, err
	}
	var bodyData types.V1FilesUploadUrlsCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1FilesUploadUrlsCreateResponse{}, err
	}
	return bodyData, nil

}
