package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	sdkgoconfig "github.com/kapetacom/sdk-go-config"
	"github.com/kapetacom/sdk-go-config/providers"
)

const (
	serviceType = "rest"
)

type RestClient struct {
	BaseURL      string
	resourceName string
	ready        bool
	mu           sync.Mutex
}

// NewRestClient initializes a new RestClient, use autoInit to automatically initialize the client when the configuration is ready.
func NewRestClient(resourceName string, autoInit bool) *RestClient {
	client := &RestClient{
		resourceName: resourceName,
	}

	if autoInit {
		sdkgoconfig.CONFIG.OnReady(func(config providers.ConfigProvider) {
			client.init(config)
		})
	}

	return client
}

// WithConfigProvider initializes the RestClient with a specific ConfigProvider.
func (c *RestClient) WithConfigProvider(config providers.ConfigProvider) *RestClient {
	c.init(config)
	return c
}

// init initializes the RestClient with the provided ConfigProvider.
func (c *RestClient) init(provider providers.ConfigProvider) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.ready {
		panic("Client already initialized")
	}

	service, err := provider.GetServiceAddress(c.resourceName, serviceType)
	if err != nil {
		panic(fmt.Sprintf("Error getting service address for %s: %s", c.resourceName, err))
	}

	c.BaseURL = strings.ToLower(service)

	c.BaseURL = strings.TrimSuffix(c.BaseURL, "/")

	log.Printf("REST client ready for %s --> %s\n", c.resourceName, c.BaseURL)
	c.ready = true
}

// ResolveURL resolves the path to a full URL by prepending the BaseURL.
func (c *RestClient) ResolveURL(path string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", c.BaseURL, fmt.Sprintf(path, args...))
}

func QueryParameterRequestModifier(queryParams any) func(req *http.Request) {
	return func(req *http.Request) {
		params, err := StructToQueryParams(queryParams)
		if err != nil {
			panic(fmt.Errorf("error creating query parameters: %s", err))
		}
		req.URL.RawQuery += params
	}
}

// GET performs a GET request to the specified URL. The requestModifier can be used to modify the request before it is sent.
// Example:
//
//	response, err := client.GET(client.ResolveURL("/api/v1/users/%s", userID), func(req *http.Request) {
//		req.Header.Set("Authorization", "Bearer "+token)
//	})
func (c *RestClient) GET(url string, requestModifier ...func(req *http.Request)) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for _, modifier := range requestModifier {
		modifier(req)
	}
	return http.DefaultClient.Do(req)
}

// DELETE performs a DELETE request to the specified URL. The requestModifier can be used to modify the request before it is sent.
// Example:
//
//	response, err := client.DELETE(client.ResolveURL("/api/v1/users/%s", userID), func(req *http.Request) {
//		req.Header.Set("Authorization", "Bearer "+token)
//	})
func (c *RestClient) DELETE(url string, requestModifier ...func(req *http.Request)) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	for _, modifier := range requestModifier {
		modifier(req)
	}
	return http.DefaultClient.Do(req)
}

// PUT performs a PUT request to the specified URL. The requestModifier can be used to modify the request before it is sent.
// Example:
//
//	response, err := client.PUT(client.ResolveURL("/api/v1/users/%s", userID), user, func(req *http.Request) {
//		req.Header.Set("Authorization", "Bearer "+token)
//	})
func (c *RestClient) PUT(url string, body any, requestModifier ...func(req *http.Request)) (*http.Response, error) {
	bodyData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for _, modifier := range requestModifier {
		modifier(req)
	}
	return http.DefaultClient.Do(req)
}

// POST performs a POST request to the specified URL. The requestModifier can be used to modify the request before it is sent.
// Example:
//
//	response, err := client.POST(client.ResolveURL("/api/v1/users/%s", userID), user, func(req *http.Request) {
//		req.Header.Set("Authorization", "Bearer "+token)
//	})
func (c *RestClient) POST(url string, body any, requestModifier ...func(req *http.Request)) (*http.Response, error) {
	bodyData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for _, modifier := range requestModifier {
		modifier(req)
	}
	return http.DefaultClient.Do(req)
}

// PATCH performs a PATCH request to the specified URL. The requestModifier can be used to modify the request before it is sent.
// Example:
//
//	response, err := client.PATCH(client.ResolveURL("/api/v1/users/%s", userID), user, func(req *http.Request) {
//		req.Header.Set("Authorization", "Bearer "+token)
//	})
func (c *RestClient) PATCH(url string, body any, requestModifier ...func(req *http.Request)) (*http.Response, error) {
	bodyData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	for _, modifier := range requestModifier {
		modifier(req)
	}
	return http.DefaultClient.Do(req)
}
