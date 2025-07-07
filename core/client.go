package core

import (
	http "net/http"
	url "net/url"
	strings "strings"
)

type CoreClient struct {
	BaseURL    map[string]string
	HttpClient *http.Client
	Auth       map[string]AuthProvider
	Modifiers  []RequestModifier
}
type RequestModifier = func(req *http.Request) error

const defaultServiceName = "__default_service__"

func DefaultBaseURL(baseURL string) map[string]string {
	return map[string]string{defaultServiceName: baseURL}
}

func NewCoreClient(baseURL map[string]string) *CoreClient {
	client := CoreClient{
		BaseURL:    baseURL,
		HttpClient: http.DefaultClient,
		Auth:       map[string]AuthProvider{},
	}
	return &client
}

func (c *CoreClient) AddAuth(request *http.Request, authNames ...string) error {
	for _, authName := range authNames {
		provider, exists := c.Auth[authName]
		if !exists {
			continue
		}
		err := provider.Apply(request)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *CoreClient) BuildURLStr(path string, serviceName ...string) string {
	// Use the provided serviceName or the default one
	name := defaultServiceName
	if len(serviceName) > 0 && serviceName[0] != "" {
		name = serviceName[0]
	}

	base, _ := c.BaseURL[name]

	// Trim trailing slash from base URL
	base = strings.TrimRight(base, "/")
	// Trim leading slash from path
	path = strings.TrimLeft(path, "/")

	return base + "/" + path
}

func (c *CoreClient) BuildURL(path string, serviceName ...string) (*url.URL, error) {
	return url.Parse(c.BuildURLStr(path, serviceName...))
}

func (c *CoreClient) ApplyModifiers(req *http.Request, modifiers []RequestModifier) error {
	// apply client-level modifier
	for _, clientMod := range c.Modifiers {
		if err := clientMod(req); err != nil {
			return err
		}
	}

	// apply req-level modifiers
	for _, reqMod := range modifiers {
		if err := reqMod(req); err != nil {
			return err
		}
	}

	return nil
}
