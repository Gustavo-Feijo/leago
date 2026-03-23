package internal

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Gustavo-Feijo/leago/ratelimit"
	nooprl "github.com/Gustavo-Feijo/leago/ratelimit/noop"
)

const (
	apiURLFormat = "https://%s.api.riotgames.com%s"
)

// Client is the shared HTTP transport used by all leago API clients.
// It holds the Riot API key, the routing prefix (platform or region), ratelimiter, logger and http doer.
type Client struct {
	HTTP        Doer
	limiter     ratelimit.RateLimiter
	Logger      *slog.Logger
	routePrefix string
	apiKey      string
}

type ClientOption func(*Client)

// WithLimiter overrides the default no-op ratelimiter.
func WithLimiter(l ratelimit.RateLimiter) ClientOption {
	return func(c *Client) { c.limiter = l }
}

// WithHTTP overrides the default http.DefaultClient with a custom doer.
func WithHTTP(d Doer) ClientOption {
	return func(c *Client) { c.HTTP = d }
}

// WithLogger sets the structured logger for request diagnostics.
// The default discards the handling.
func WithLogger(l *slog.Logger) ClientOption {
	return func(c *Client) { c.Logger = l }
}

// NewHTTPClient creates a Client scoped to a single Riot routing value (e.g. "euw1", "europe").
// By default it uses http.DefaultClient, a no-op rate limiter and a discarding logger.
func NewHTTPClient(route, apiKey string, opts ...ClientOption) *Client {
	c := &Client{
		HTTP:        http.DefaultClient,
		limiter:     nooprl.NewNoopLimiter(),
		Logger:      slog.New(slog.NewTextHandler(io.Discard, nil)),
		routePrefix: route,
		apiKey:      apiKey,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// GetURL constructs the full Riot API URL for the given endpoint path.
func (c *Client) GetURL(endpoint string) string {
	return fmt.Sprintf(apiURLFormat, c.routePrefix, endpoint)
}
