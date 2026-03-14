package lor

import (
	"github.com/Gustavo-Feijo/leago/api/lor/matches"
	"github.com/Gustavo-Feijo/leago/api/lor/ranked"
	"github.com/Gustavo-Feijo/leago/internal"
)

type RegionClient struct {
	Matches *matches.RegionClient
	Ranked  *ranked.RegionClient
}

func NewRegionClient(baseClient *internal.Client) *RegionClient {
	c := &RegionClient{
		Matches: matches.NewRegionClient(baseClient),
		Ranked:  ranked.NewRegionClient(baseClient),
	}
	return c
}
