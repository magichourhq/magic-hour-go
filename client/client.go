package client

import (
	sdkcore "github.com/magichourhq/magic-hour-go/core"
	v1 "github.com/magichourhq/magic-hour-go/resources/v1"
)

type Client struct {
	coreClient *sdkcore.CoreClient
	V1         *v1.Client
}

// Instantiate a new API client
func NewClient(builders ...func(*sdkcore.CoreClient)) *Client {
	defaultEnv := Environment
	coreClient := sdkcore.NewCoreClient(string(defaultEnv))
	for _, b := range builders {
		b(coreClient)
	}

	client := Client{
		coreClient: coreClient,
		V1:         v1.NewClient(coreClient),
	}

	return &client
}
