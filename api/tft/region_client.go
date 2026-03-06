package tft

import (
	"log/slog"

	"github.com/Gustavo-Feijo/leago/api/tft/match"
	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/regions"
)

type RegionClient struct {
	Match *match.RegionClient
}

func NewRegionClient(client internal.Doer, logger *slog.Logger, region regions.Region, apiKey string) *RegionClient {
	baseClient := internal.NewHTTPClient(client, logger, string(region), apiKey)
	c := &RegionClient{
		Match: match.NewRegionClient(baseClient),
	}
	return c
}
