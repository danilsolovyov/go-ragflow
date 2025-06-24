package options

import (
	"net/http"
	"time"
)

const (
	ClientDefaultTimeout = 30 * time.Second
)

type ClientOptions struct {
	Scheme    string
	Host      string
	Transport http.RoundTripper
	Timeout   time.Duration
	APIKey    string
}

func DefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		Scheme:    "http",
		Host:      "localhost:8080",
		Transport: http.DefaultTransport,
		Timeout:   ClientDefaultTimeout,
	}
}

func (o *ClientOptions) Merge(other *ClientOptions) *ClientOptions {
	if other == nil {
		return o
	}

	if other.Scheme != "" {
		o.Scheme = other.Scheme
	}
	if other.Host != "" {
		o.Host = other.Host
	}
	if other.Transport != nil {
		o.Transport = other.Transport
	}
	if other.Timeout > 0 {
		o.Timeout = other.Timeout
	}
	if other.APIKey != "" {
		o.APIKey = other.APIKey
	}

	return o
}

func (o *ClientOptions) SetScheme(scheme string) {
	o.Scheme = scheme
}

func (o *ClientOptions) SetHost(host string) {
	o.Host = host
}

func (o *ClientOptions) SetTransport(transport http.RoundTripper) {
	o.Transport = transport
}

func (o *ClientOptions) SetTimeout(timeout time.Duration) {
	o.Timeout = timeout
}

func (o *ClientOptions) SetAPIKey(apiKey string) {
	o.APIKey = apiKey
}
