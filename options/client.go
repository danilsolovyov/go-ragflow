package options

import (
	"net/http"
	"time"
)

const (
	defaultScheme  = "http"
	defaultHost    = "localhost:8080"
	defaultTimeout = 30 * time.Second
)

type ClientOptions struct {
	Scheme    string
	Host      string
	Transport http.RoundTripper
	Timeout   time.Duration
}

func DefaultClientOptions() *ClientOptions {
	return &ClientOptions{
		Scheme:    "http",
		Host:      "localhost:8080",
		Transport: http.DefaultTransport,
		Timeout:   30 * time.Second,
	}
}

func (co *ClientOptions) SetScheme(scheme string) {
	co.Scheme = scheme
}

func (co *ClientOptions) SetHost(host string) {
	co.Host = host
}

func (co *ClientOptions) SetTransport(transport http.RoundTripper) {
	co.Transport = transport
}

func (co *ClientOptions) SetTimeout(timeout time.Duration) {
	co.Timeout = timeout
}
