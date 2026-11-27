package estimate

import (
	http "net/http"

	sdkcore "github.com/magichourhq/magic-hour-go/core"
	image "github.com/magichourhq/magic-hour-go/resources/v1/estimate/image"
)

type Client struct {
	coreClient *sdkcore.CoreClient
	Image      *image.Client
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient: coreClient,
		Image:      image.NewClient(coreClient),
	}

	return &client
}
