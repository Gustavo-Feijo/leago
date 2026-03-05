package riot

import (
	"log/slog"

	"github.com/Gustavo-Feijo/leago/api/riot/account"
	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/regions"
)

type RegionClient struct {
	Account *account.RegionClient
}

func NewRegionClient(client internal.Doer, logger *slog.Logger, region regions.Region, apiKey string) *RegionClient {
	baseClient := internal.NewHTTPClient(client, logger, string(region), apiKey)
	c := &RegionClient{
		Account: account.NewRegionClient(baseClient),
	}
	return c
}
