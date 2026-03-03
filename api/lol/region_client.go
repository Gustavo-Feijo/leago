package lol

import (
	"log/slog"

	"github.com/Gustavo-Feijo/leago/api/lol/matches"
	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/regions"
)

type RegionClient struct {
	Matches *matches.RegionClient
}

func NewRegionClient(client internal.Doer, logger *slog.Logger, region regions.Region, apiKey string) *RegionClient {
	baseClient := internal.NewHttpClient(client, logger, string(region), apiKey)
	c := &RegionClient{
		Matches: matches.NewRegionClient(baseClient),
	}
	return c
}
