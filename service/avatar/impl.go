package avatar

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/SourceGlobalCDN/avatar-proxy/pkg/env"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/log"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/request"
)

type Impl struct {
	client  *request.Client
	baseUrl *url.URL
}

var _ Factory = &Impl{}

func NewImpl() *Impl {
	client := request.NewClient()
	client.SetBaseUrl(env.ProxyConfig.Remote)
	client.SetHeader("User-Agent", env.ProxyConfig.UserAgent)

	base, _ := url.Parse(env.ProxyConfig.Remote)

	return &Impl{
		client:  client,
		baseUrl: base,
	}
}

func (i *Impl) GetAvatar(hash string, option Payload) (*io.ReadCloser, int, error) {
	u, err := i.baseUrl.Parse(hash)
	if err != nil {
		return nil, 0, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, 0, err
	}

	q := req.URL.Query()
	if option.Size > 0 {
		q.Add("s", strconv.Itoa(option.Size))
	}
	if option.DefaultImg != "" {
		q.Add("d", option.DefaultImg)
	}
	if option.ForceDefault {
		q.Add("f", "y")
	}
	if option.Rating != "" {
		q.Add("r", string(option.Rating))
	}

	req.URL.RawQuery = q.Encode()

	log.Log().Debugf("Request: %s", req.URL.String())

	res, err := i.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	if res.StatusCode != 200 {
		return nil, 0, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	length, _ := strconv.Atoi(res.Header.Get("Content-Length"))

	return &res.Body, length, nil
}
