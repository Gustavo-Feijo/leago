package lor

import (
	"leago/api/lor/matches"
	"leago/internal"
	"leago/regions"
	"log/slog"
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
