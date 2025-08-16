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
// Generates a list of pre-signed upload URLs for the assets required. This API is only necessary if you want to upload to Magic Hour's storage. Refer to the [Input Files Guide](/integration/input-files) for more details.
//
// The response array will match the order of items in the request body.
//
// **Valid file extensions per asset type**:
// - video: mp4, m4v, mov, webm
// - audio: mp3, mpeg, wav, aac, aiff, flac
// - image: png, jpg, jpeg, webp, avif, jp2, tiff, bmp
//
// > Note: `gif` is only supported for face swap API `video_file_path` field.
//
// Once you receive an upload URL, send a `PUT` request to upload the file directly.
//
// Example:
//
// ```
//
//	curl -X PUT --data '@/path/to/file/video.mp4' \
//	  https://videos.magichour.ai/api-assets/id/video.mp4?<auth params from the API response>
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
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1FilesUploadUrlsCreateResponse{}, err
	}

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
