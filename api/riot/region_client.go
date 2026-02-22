package riot

import (
	"leago/api/riot/account"
	"leago/internal"
	"leago/regions"
	"log/slog"
)

type RegionClient struct {
	Account *account.RegionClient
}

func NewRegionClient(client internal.Doer, logger *slog.Logger, region regions.Region, apiKey string) *RegionClient {
	baseClient := internal.NewHttpClient(client, logger, string(region), apiKey)
	c := &RegionClient{
		Account: account.NewRegionClient(baseClient),
	}
	return c
}
