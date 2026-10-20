package ai_ugc_ad

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

// AI UGC Ad Generator
//
// Create a spoken UGC video from a prepared image and script. Returns a project ID; poll GET /v1/video-projects/{id} for status and downloads. To compose an image, generate the image separately with the Image Editor API and supply its download URL.
//
// POST /v1/ai-ugc-ad
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiUgcAdCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "ai-ugc-ad")
	if err != nil {
		return types.V1AiUgcAdCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AiUgcAdCreateBody{
		AdditionalProperties: request.AdditionalProperties,
		Model:                request.Model,
		Resolution:           request.Resolution,
		Assets:               request.Assets,
		EndSeconds:           request.EndSeconds,
		Script:               request.Script,
	})
	if err != nil {
		return types.V1AiUgcAdCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiUgcAdCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1AiUgcAdCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiUgcAdCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiUgcAdCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiUgcAdCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiUgcAdCreateResponse{}, err
	}
	var bodyData types.V1AiUgcAdCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiUgcAdCreateResponse{}, err
	}
	return bodyData, nil

}
