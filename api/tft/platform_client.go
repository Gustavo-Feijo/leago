package tft

import (
	"log/slog"

	"github.com/Gustavo-Feijo/leago/api/tft/league"
	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/regions"
)

type PlatformClient struct {
	League *league.PlatformClient
}

func NewPlatformClient(client internal.Doer, logger *slog.Logger, region regions.Platform, apiKey string) *PlatformClient {
	baseClient := internal.NewHTTPClient(client, logger, string(region), apiKey)
	c := &PlatformClient{
		League: league.NewPlatformClient(baseClient),
	}
	return c
}
