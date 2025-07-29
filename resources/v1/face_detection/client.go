package face_detection

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

// Get face detection details
//
// Get the details of a face detection task.
//
// GET /v1/face-detection/{id}
func (c *Client) Get(request GetRequest, reqModifiers ...RequestModifier) (types.V1FaceDetectionGetResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "face-detection/" + sdkcore.FmtStringParam(request.Id))
	if err != nil {
		return types.V1FaceDetectionGetResponse{}, err
	}

	// Init request
	req, err := http.NewRequest("GET", targetUrl.String(), nil)
	if err != nil {
		return types.V1FaceDetectionGetResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1FaceDetectionGetResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1FaceDetectionGetResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1FaceDetectionGetResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1FaceDetectionGetResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1FaceDetectionGetResponse{}, err
	}
	var bodyData types.V1FaceDetectionGetResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1FaceDetectionGetResponse{}, err
	}
	return bodyData, nil

}

// Face Detection
//
// Detect faces in an image or video.
//
// Note: Face detection is free to use for the near future. Pricing may change in the future.
//
// POST /v1/face-detection
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1FaceDetectionCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "face-detection")
	if err != nil {
		return types.V1FaceDetectionCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1FaceDetectionCreateBody{
		ConfidenceScore: request.ConfidenceScore,
		Assets:          request.Assets,
	})
	if err != nil {
		return types.V1FaceDetectionCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1FaceDetectionCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1FaceDetectionCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1FaceDetectionCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1FaceDetectionCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1FaceDetectionCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1FaceDetectionCreateResponse{}, err
	}
	var bodyData types.V1FaceDetectionCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1FaceDetectionCreateResponse{}, err
	}
	return bodyData, nil

}
