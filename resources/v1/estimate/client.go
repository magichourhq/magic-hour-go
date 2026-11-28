package estimate

import (
	http "net/http"

	sdkcore "github.com/magichourhq/magic-hour-go/core"
	audio "github.com/magichourhq/magic-hour-go/resources/v1/estimate/audio"
	image "github.com/magichourhq/magic-hour-go/resources/v1/estimate/image"
	video "github.com/magichourhq/magic-hour-go/resources/v1/estimate/video"
)

type Client struct {
	coreClient *sdkcore.CoreClient
	Audio      *audio.Client
	Image      *image.Client
	Video      *video.Client
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient: coreClient,
		Audio:      audio.NewClient(coreClient),
		Image:      image.NewClient(coreClient),
		Video:      video.NewClient(coreClient),
	}

	return &client
}
