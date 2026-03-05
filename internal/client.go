package internal

import (
	"fmt"
	"log/slog"
)

const (
	apiURLFormat = "https://%s.api.riotgames.com%s"
)

type Client struct {
	HTTP        Doer
	Logger      *slog.Logger
	routePrefix string
	apiKey      string
}

func NewHTTPClient(client Doer, logger *slog.Logger, route, apiKey string) *Client {
	c := &Client{
		HTTP:        client,
		Logger:      logger,
		routePrefix: route,
		apiKey:      apiKey,
	}

	return c
}

func (c *Client) GetURL(endpoint string) string {
	return fmt.Sprintf(apiURLFormat, c.routePrefix, endpoint)
}
