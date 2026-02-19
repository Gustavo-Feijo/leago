package internal

import (
	"fmt"
)

const (
	apiURLFormat = "https://%s.api.riotgames.com%s"
)

type Client struct {
	Http        Doer
	routePrefix string
	ApiKey      string
}

func NewHttpClient(client Doer, route, apiKey string) *Client {
	c := &Client{
		Http:        client,
		routePrefix: route,
		ApiKey:      apiKey,
	}

	return c
}

func (c *Client) GetURL(endpoint string) string {
	return fmt.Sprintf(apiURLFormat, c.routePrefix, endpoint)
}
