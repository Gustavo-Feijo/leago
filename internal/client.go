package internal

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Gustavo-Feijo/leago/ratelimit"
	nooprl "github.com/Gustavo-Feijo/leago/ratelimit/noop"
)

const (
	apiURLFormat = "https://%s.api.riotgames.com%s"
)

type Client struct {
	HTTP        Doer
	limiter     ratelimit.RateLimiter
	Logger      *slog.Logger
	routePrefix string
	apiKey      string
}

type ClientOption func(*Client)

func WithLimiter(l ratelimit.RateLimiter) ClientOption {
	return func(c *Client) { c.limiter = l }
}

func WithHTTP(d Doer) ClientOption {
	return func(c *Client) { c.HTTP = d }
}

func WithLogger(l *slog.Logger) ClientOption {
	return func(c *Client) { c.Logger = l }
}

func NewHTTPClient(route, apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		HTTP:        http.DefaultClient,
		limiter:     nooprl.NewNoopLimiter(),
		Logger:      slog.New(slog.DiscardHandler),
		routePrefix: route,
		apiKey:      apiKey,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) GetURL(endpoint string) string {
	return fmt.Sprintf(apiURLFormat, c.routePrefix, endpoint)
}
