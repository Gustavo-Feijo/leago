package riot

import (
	"leago/api/riot/account"
	"leago/internal"
	"leago/regions"
)

type RegionClient struct {
	Account *account.RegionClient
}

func NewRegionClient(client internal.Doer, region regions.Region, apiKey string) *RegionClient {
	baseClient := internal.NewHttpClient(client, string(region), apiKey)
	c := &RegionClient{
		Account: account.NewRegionClient(baseClient),
	}
	return c
}
