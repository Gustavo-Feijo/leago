package riot

import (
	"github.com/Gustavo-Feijo/leago/api/riot/account"
	"github.com/Gustavo-Feijo/leago/internal"
)

// RegionClient groups the Riot APIs that are scoped to a routing region (e.g. EUROPE, AMERICAS).
type RegionClient struct {
	Account *account.RegionClient
}

func NewRegionClient(baseClient *internal.Client) *RegionClient {
	c := &RegionClient{
		Account: account.NewRegionClient(baseClient),
	}
	return c
}
