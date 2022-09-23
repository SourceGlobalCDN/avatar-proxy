package request

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	asserts := assert.New(t)

	client := NewClient()
	asserts.NotNil(client)
}

func TestClient_SetHeader(t *testing.T) {
	asserts := assert.New(t)

	client := NewClient()
	client.SetHeader("test", "test")
	asserts.Equal("test", client.headers.Get("test"))
}

func TestClient_SetHeaders(t *testing.T) {
	asserts := assert.New(t)

	client := NewClient()
	client.SetHeaders(http.Header{"test": {"test"}})
	asserts.Equal("test", client.headers.Get("test"))
}

func TestClient_SetBaseUrl(t *testing.T) {
	asserts := assert.New(t)

	client := NewClient()
	client.SetBaseUrl("https://test.com")
	asserts.Equal("https://test.com", client.baseUrl.String())
}

func TestClient_GetBaseUrl(t *testing.T) {
	asserts := assert.New(t)

	client := NewClient()
	client.SetBaseUrl("https://test.com")
	asserts.Equal("https://test.com", client.GetBaseUrl().String())
}

func TestClient_Do(t *testing.T) {
	asserts := assert.New(t)

	client := NewClient()
	req, err := http.NewRequest("GET", "https://example.com", nil)
	asserts.NoError(err)

	res, err := client.Do(req)
	asserts.NoError(err)
	asserts.NotNil(res)
}

func TestClient_Get(t *testing.T) {
	asserts := assert.New(t)

	client := NewClient()
	res, err := client.Get("https://example.com")
	asserts.NoError(err)
	asserts.NotNil(res)
}

func TestClient_Post(t *testing.T) {
	asserts := assert.New(t)

	client := NewClient()
	res, err := client.Post("https://example.com", "text/plain", nil)
	asserts.NoError(err)
	asserts.NotNil(res)
}
