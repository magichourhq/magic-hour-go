package client

import (
	sdkcore "github.com/magichourhq/magic-hour-go/core"
	http "net/http"
)

// Provide your own http.Client to be used for all requests
func WithHTTPClient(httpClient *http.Client) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.HttpClient = httpClient
	}
}

func WithEnv(env Env) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.BaseURL = string(env)
	}
}

// Provide non-default baseURL for all requests
func WithBaseURL(baseURL string) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.BaseURL = baseURL
	}
}

type RequestModifier = func(req *http.Request) error

// Provide modifiers to be applied to all client requests
func WithModifiers(modifiers ...RequestModifier) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.Modifiers = append(c.Modifiers, modifiers...)
	}
}

func WithBearerAuth(token string) func(*sdkcore.CoreClient) {
	return func(c *sdkcore.CoreClient) {
		c.Auth["bearerAuth"] = func(request *http.Request) {
			request.Header.Add("Authorization", "Bearer "+token)
		}
	}
}
