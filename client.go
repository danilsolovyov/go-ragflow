package goragflow

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/danilsolovyov/go-ragflow/options"
	"github.com/danilsolovyov/go-ragflow/parameters"
)

const (
	apiPath = "/api/v1"
)

type Client struct {
	httpClient *http.Client
	scheme     string
	host       string
	timeout    time.Duration
}

// authTransport is a RoundTripper that adds the API key to the request.
type authTransport struct {
	defaultRoundTripper http.RoundTripper
	apiKey              string
}

// RoundTrip implements the RoundTripper interface.
func (a *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+a.apiKey)
	req.Header.Set("Content-Type", "application/json")
	return a.defaultRoundTripper.RoundTrip(req)
}

// NewClient creates a new RagFlow client.
func NewClient(opts *options.ClientOptions) *Client {
	opts = options.DefaultClientOptions().Merge(opts)

	client := &Client{
		httpClient: &http.Client{
			Transport: opts.Transport,
			Timeout:   opts.Timeout,
		},
		scheme: opts.Scheme,
		host:   opts.Host,
	}
	authTransport := &authTransport{
		defaultRoundTripper: client.httpClient.Transport,
		apiKey:              opts.APIKey,
	}

	client.httpClient.Transport = authTransport

	return client
}

func (c *Client) do(ctx context.Context, method, path string, data any, params ...parameters.Parameter) error {
	u := &url.URL{
		Scheme: c.scheme,
		Host:   c.host,
		Path:   apiPath + path,
	}

	err := parameters.ApplyURL(u, params...)
	if err != nil {
		return err
	}
	var reqBody io.Reader

	// TODO: add support for files multipart/form-data
	body := parameters.CreateBody(params...)

	if len(body) > 0 {
		reqBody, err = ioReaderFromStruct(body)
		if err != nil {
			return err
		}
	} else {
		reqBody = http.NoBody
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reqBody)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData := Response{
		Data: data,
	}

	if err = json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unknown error, unexpected status code: %d", resp.StatusCode)
	}

	if respData.Code != 0 {
		return fmt.Errorf("unknown error, unexpected request code: %w", ErrorResponse(respData))
	}

	return nil
}
func (r *Client) Close() {
	if r.httpClient != nil {
		r.httpClient.CloseIdleConnections()
	}
}
