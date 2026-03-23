package lor

import (
	"github.com/Gustavo-Feijo/leago/api/lor/matches"
	"github.com/Gustavo-Feijo/leago/api/lor/ranked"
	"github.com/Gustavo-Feijo/leago/api/lor/status"
	"github.com/Gustavo-Feijo/leago/internal"
)

// RegionClient groups the Legends Of Runeterra APIs that are scoped to a routing region (e.g. EUROPE, AMERICAS).
type RegionClient struct {
	Matches *matches.RegionClient
	Ranked  *ranked.RegionClient
	Status  *status.RegionClient
}

func NewRegionClient(baseClient *internal.Client) *RegionClient {
	c := &RegionClient{
		Matches: matches.NewRegionClient(baseClient),
		Ranked:  ranked.NewRegionClient(baseClient),
		Status:  status.NewRegionClient(baseClient),
	}
	return c
}
