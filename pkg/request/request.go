package request

import (
	"bytes"
	"net/http"
	"net/url"
	"time"

	"github.com/SourceGlobalCDN/avatar-proxy/pkg/env"
)

type Client struct {
	http.ServeMux

	client *http.Client

	headers http.Header
	baseUrl *url.URL
}

func NewClient() *Client {
	client := &Client{
		client:  http.DefaultClient,
		headers: make(http.Header),
		baseUrl: nil,
	}

	client.client.Timeout = time.Duration(env.ProxyConfig.Timeout) * time.Second

	return client
}

func (c *Client) SetHeader(key, value string) {
	c.headers.Set(key, value)
}

func (c *Client) SetHeaders(headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			c.SetHeader(key, value)
		}
	}
}

func (c *Client) SetBaseUrl(baseUrl string) {
	if base, err := url.Parse(baseUrl); err == nil {
		c.baseUrl = base
	}
}

func (c *Client) GetBaseUrl() *url.URL {
	return c.baseUrl
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	for key, values := range c.headers {
		for _, value := range values {
			req.Header.Set(key, value)
		}
	}

	req.URL = c.baseUrl.ResolveReference(req.URL)

	return c.client.Do(req)
}

func (c *Client) Get(path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Post(path string, bodyType string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", bodyType)

	return c.Do(req)
}
