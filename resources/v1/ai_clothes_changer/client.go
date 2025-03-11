package ai_clothes_changer

import (
	bytes "bytes"
	json "encoding/json"
	sdkcore "github.com/magichourhq/magic-hour-go/core"
	types "github.com/magichourhq/magic-hour-go/types"
	io "io"
	http "net/http"
	url "net/url"
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
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.PostV1AiClothesChangerResponse, error) {
	// URL formatting
	joinedUrl, err := url.JoinPath(c.coreClient.BaseURL, "/v1/"+"ai-clothes-changer")
	if err != nil {
		return types.PostV1AiClothesChangerResponse{}, err
	}
	targetUrl, err := url.Parse(joinedUrl)
	if err != nil {
		return types.PostV1AiClothesChangerResponse{}, err
	}

	// Prep body
	reqBody, err := json.Marshal(types.PostV1AiClothesChangerBody{
		Name:   request.Name,
		Assets: request.Assets})
	if err != nil {
		return types.PostV1AiClothesChangerResponse{}, err
	}
	reqBodyBuf := bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.PostV1AiClothesChangerResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	c.coreClient.AddAuth([]string{"bearerAuth"}, req)

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.PostV1AiClothesChangerResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.PostV1AiClothesChangerResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.PostV1AiClothesChangerResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.PostV1AiClothesChangerResponse{}, err
	}
	var bodyData types.PostV1AiClothesChangerResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.PostV1AiClothesChangerResponse{}, err
	}
	return bodyData, nil

}
