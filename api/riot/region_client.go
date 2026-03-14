package riot

import (
	"github.com/Gustavo-Feijo/leago/api/riot/account"
	"github.com/Gustavo-Feijo/leago/internal"
)

type RegionClient struct {
	Account *account.RegionClient
}

func NewRegionClient(baseClient *internal.Client) *RegionClient {
	c := &RegionClient{
		Account: account.NewRegionClient(baseClient),
	}
	return c
}
