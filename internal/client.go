package internal

import (
	"fmt"
	"log/slog"
)

const (
	apiURLFormat = "https://%s.api.riotgames.com%s"
)

type Client struct {
	Http        Doer
	Logger      *slog.Logger
	routePrefix string
	apiKey      string
}

func NewHttpClient(client Doer, logger *slog.Logger, route, apiKey string) *Client {
	c := &Client{
		Http:        client,
		Logger:      logger,
		routePrefix: route,
		apiKey:      apiKey,
	}

	return c
}

func (c *Client) GetURL(endpoint string) string {
	return fmt.Sprintf(apiURLFormat, c.routePrefix, endpoint)
}
