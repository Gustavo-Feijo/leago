package lor

import (
	"leago/api/lor/matches"
	"leago/api/lor/ranked"
	"leago/internal"
	"leago/regions"
	"log/slog"
)

type RegionClient struct {
	Matches *matches.RegionClient
	Ranked  *ranked.RegionClient
}

func NewRegionClient(client internal.Doer, logger *slog.Logger, region regions.Region, apiKey string) *RegionClient {
	baseClient := internal.NewHttpClient(client, logger, string(region), apiKey)
	c := &RegionClient{
		Matches: matches.NewRegionClient(baseClient),
		Ranked:  ranked.NewRegionClient(baseClient),
	}
	return c
}
