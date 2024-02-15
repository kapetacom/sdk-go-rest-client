package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	config "github.com/kapetacom/sdk-go-config"
	"github.com/stretchr/testify/assert"
)

func TestNewRestClient(t *testing.T) {
	t.Run("should initialize a new RestClient", func(t *testing.T) {
		client := NewRestClient("resource", false)
		assert.NotNil(t, client)
	})
	t.Run("should initialize a new RestClient and autoInit", func(t *testing.T) {
		client := NewRestClient("resource", true)
		assert.NotNil(t, client)
	})
	t.Run("should initialize a new RestClient with a specific ConfigProvider", func(t *testing.T) {
		mock := &config.ConfigProviderMock{
			GetServiceAddressFunc: func(serviceName string, portType string) (string, error) {
				return "", nil
			},
		}
		client := NewRestClient("resource", false).WithConfigProvider(mock)
		assert.NotNil(t, client)
	})
	t.Run("should be able to call get on client without error", func(t *testing.T) {
		called := false
		srv := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet {
					called = true
				}
			}),
		)

		mock := &config.ConfigProviderMock{
			GetServiceAddressFunc: func(serviceName string, portType string) (string, error) {
				return "", nil
			},
		}
		client := NewRestClient("resource", false).WithConfigProvider(mock)
		_, err := client.GET(srv.URL)
		assert.Nil(t, err)
		assert.True(t, called)
	})
	t.Run("should be able to call post on client without error", func(t *testing.T) {
		called := false
		srv := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPost {
					called = true
				}
			}),
		)

		mock := &config.ConfigProviderMock{
			GetServiceAddressFunc: func(serviceName string, portType string) (string, error) {
				return "", nil
			},
		}
		client := NewRestClient("resource", false).WithConfigProvider(mock)
		_, err := client.POST(srv.URL, nil)
		assert.Nil(t, err)
		assert.True(t, called)
	})
	t.Run("should be able to call put on client without error", func(t *testing.T) {
		called := false
		srv := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPut {
					called = true
				}
			}),
		)

		mock := &config.ConfigProviderMock{
			GetServiceAddressFunc: func(serviceName string, portType string) (string, error) {
				return "", nil
			},
		}
		client := NewRestClient("resource", false).WithConfigProvider(mock)
		_, err := client.PUT(srv.URL, nil)
		assert.Nil(t, err)
		assert.True(t, called)
	})
	t.Run("should be able to call delete on client without error", func(t *testing.T) {
		called := false
		srv := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodDelete {
					called = true
				}
			}),
		)

		mock := &config.ConfigProviderMock{
			GetServiceAddressFunc: func(serviceName string, portType string) (string, error) {
				return "", nil
			},
		}
		client := NewRestClient("resource", false).WithConfigProvider(mock)
		_, err := client.DELETE(srv.URL)
		assert.Nil(t, err)
		assert.True(t, called)
	})
	t.Run("should be able to call patch on client without error", func(t *testing.T) {
		called := false
		srv := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodPatch {
					called = true
				}
			}),
		)

		mock := &config.ConfigProviderMock{
			GetServiceAddressFunc: func(serviceName string, portType string) (string, error) {
				return "", nil
			},
		}
		client := NewRestClient("resource", false).WithConfigProvider(mock)
		_, err := client.PATCH(srv.URL, nil)
		assert.Nil(t, err)
		assert.True(t, called)
	})

	t.Run("should be able to call get on client with request modifier without error", func(t *testing.T) {
		called := false
		srv := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet {
					if r.Header.Get("Authorization") == "Bearer token" {
						called = true
					}
				}
			}),
		)

		mock := &config.ConfigProviderMock{
			GetServiceAddressFunc: func(serviceName string, portType string) (string, error) {
				return "", nil
			},
		}
		client := NewRestClient("resource", false).WithConfigProvider(mock)
		_, err := client.GET(srv.URL, func(req *http.Request) {
			req.Header.Set("Authorization", "Bearer token")
		})
		assert.Nil(t, err)
		assert.True(t, called)
	})
	t.Run("should be able to call get on client witrh request modifier that adds query params without error", func(t *testing.T) {
		called := false
		type Params struct {
			Param string
		}
		srv := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodGet {
					if r.URL.Query().Get("param") == "value" {
						called = true
					}
				}
			}),
		)

		mock := &config.ConfigProviderMock{
			GetServiceAddressFunc: func(serviceName string, portType string) (string, error) {
				return "", nil
			},
		}
		client := NewRestClient("resource", false).WithConfigProvider(mock)
		_, err := client.GET(srv.URL, QueryParameterRequestModifier(Params{Param: "value"}))
		assert.Nil(t, err)
		assert.True(t, called)
	})
}
