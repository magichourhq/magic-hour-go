package ai_clothes_changer

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

// AI Clothes Changer
//
// Change outfits in photos in seconds with just a photo reference. Each photo costs 25 frames.
//
// POST /v1/ai-clothes-changer
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AiClothesChangerCreateResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"ai-clothes-changer")
	if err != nil {
		return types.V1AiClothesChangerCreateResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.V1AiClothesChangerCreateResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(types.V1AiClothesChangerCreateBody{
		Name:   request.Name,
		Assets: request.Assets})
	if err != nil {
		return types.V1AiClothesChangerCreateResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AiClothesChangerCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AiClothesChangerCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AiClothesChangerCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AiClothesChangerCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AiClothesChangerCreateResponse{}, err
	}
	var bodyData types.V1AiClothesChangerCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AiClothesChangerCreateResponse{}, err
	}
	return bodyData, nil

}
