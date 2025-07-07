package client

import (
	http "net/http"

	sdkcore "github.com/magichourhq/magic-hour-go/core"
)

type RequestModifier = func(req *http.Request) error

// Provide modifiers to be applied to all client requests
func WithModifiers(modifiers ...RequestModifier) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.Modifiers = append(c.Modifiers, modifiers...)
	}
}

// Customize baseURL using pre-defined environments
func WithEnv(env Env) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.BaseURL = sdkcore.DefaultBaseURL(env.String())
	}
}

// Provide non-default baseURL for all requests
func WithBaseURL(baseURL string) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.BaseURL = sdkcore.DefaultBaseURL(baseURL)
	}
}

// Provide your own http.Client to be used for all requests
func WithHTTPClient(httpClient *http.Client) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.HttpClient = httpClient
	}
}

func WithBearerAuth(token string) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.Auth["bearerAuth"] = sdkcore.NewAuthBearer(token)
	}
}
