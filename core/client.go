package core

import (
	http "net/http"
)

type CoreClient struct {
	BaseURL    string
	HttpClient *http.Client
	Auth       map[string]func(*http.Request)
	Modifiers  []RequestModifier
}
type RequestModifier = func(req *http.Request) error

func NewCoreClient(baseURL string) *CoreClient {
	client := CoreClient{
		BaseURL:    baseURL,
		HttpClient: http.DefaultClient,
		Auth:       map[string]func(*http.Request){},
	}
	return &client
}

func (c *CoreClient) AddAuth(authNames []string, request *http.Request) {
	for _, authName := range authNames {
		provider, exists := c.Auth[authName]
		if !exists {
			continue
		}
		provider(request)
	}
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
