package auto_subtitle_generator

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

// Auto Subtitle Generator
//
// Automatically generate subtitles for your video in multiple languages.
//
// POST /v1/auto-subtitle-generator
func (c *Client) Create(request CreateRequest, reqModifiers ...RequestModifier) (types.V1AutoSubtitleGeneratorCreateResponse, error) {
	// URL formatting
	targetUrl, err := c.coreClient.BuildURL("/v1/" + "auto-subtitle-generator")
	if err != nil {
		return types.V1AutoSubtitleGeneratorCreateResponse{}, err
	}

	// Prep body
	reqBodyBuf := &bytes.Buffer{}
	reqBody, err := json.Marshal(types.V1AutoSubtitleGeneratorCreateBody{
		Name:         request.Name,
		Assets:       request.Assets,
		EndSeconds:   request.EndSeconds,
		StartSeconds: request.StartSeconds,
		Style:        request.Style,
	})
	if err != nil {
		return types.V1AutoSubtitleGeneratorCreateResponse{}, err
	}
	reqBodyBuf = bytes.NewBuffer([]byte(reqBody))

	// Init request
	req, err := http.NewRequest("POST", targetUrl.String(), reqBodyBuf)
	if err != nil {
		return types.V1AutoSubtitleGeneratorCreateResponse{}, err
	}

	// Add headers
	req.Header.Add("x-sideko-sdk-language", "Go")
	req.Header.Add("Content-Type", "application/json")

	// Add auth
	err = c.coreClient.AddAuth(req, "bearerAuth")
	if err != nil {
		return types.V1AutoSubtitleGeneratorCreateResponse{}, err
	}

	// Add base client & request level modifiers
	if err := c.coreClient.ApplyModifiers(req, reqModifiers); err != nil {
		return types.V1AutoSubtitleGeneratorCreateResponse{}, err
	}

	// Dispatch request
	resp, err := c.coreClient.HttpClient.Do(req)
	if err != nil {
		return types.V1AutoSubtitleGeneratorCreateResponse{}, err
	}

	// Check status
	if resp.StatusCode >= 300 {
		return types.V1AutoSubtitleGeneratorCreateResponse{}, sdkcore.NewApiError(*req, *resp)
	}

	// Handle response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.V1AutoSubtitleGeneratorCreateResponse{}, err
	}
	var bodyData types.V1AutoSubtitleGeneratorCreateResponse
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		return types.V1AutoSubtitleGeneratorCreateResponse{}, err
	}
	return bodyData, nil

}
