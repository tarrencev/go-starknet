package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Provider struct {
	baseURL          string
	feederGatewayURL string
	gatewayURL       string
	c                *http.Client
	errorHandler     func(e error) error
}

// Option sets an optional setting on the Client.
type Option func(*Provider)

func NewProvider(baseURL string, opts ...Option) *Provider {
	p := &Provider{
		baseURL:          baseURL,
		feederGatewayURL: baseURL + "/feeder_gateway",
		gatewayURL:       baseURL + "/gateway",
		c:                http.DefaultClient,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// WithHTTPClient returns an Option to set the http.Client to be used.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(p *Provider) {
		p.c = httpClient
	}
}

// WithBaseURL returns an Option to set the base URL to be used.
func WithBaseURL(baseURL string) Option {
	return func(p *Provider) {
		p.baseURL = baseURL
	}
}

// WithErrorHandler returns an Option to set the error handler to be used.
func WithErrorHandler(f func(e error) error) Option {
	return func(p *Provider) {
		p.errorHandler = f
	}
}

func (p *Provider) newRequest(
	ctx context.Context, method, endpoint string, body interface{},
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, p.feederGatewayURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		req.Header.Add("Content-Type", "application/json; charset=utf")
	}
	return req, nil
}

func (p *Provider) do(req *http.Request, v interface{}) error {
	resp, err := p.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // nolint: errcheck

	if resp.StatusCode >= 299 {
		e := NewError(resp)
		if p.errorHandler != nil {
			return p.errorHandler(e)
		}
		return e
	}

	if v != nil {
		return json.NewDecoder(resp.Body).Decode(v)
	}
	return nil
}

func appendQueryValues(req *http.Request, values url.Values) {
	q := req.URL.Query()
	for k, vs := range values {
		for _, v := range vs {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()
}
