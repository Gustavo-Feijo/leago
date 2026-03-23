package lol

import (
	"github.com/Gustavo-Feijo/leago/api/lol/match"
	"github.com/Gustavo-Feijo/leago/api/lol/tournament"
	"github.com/Gustavo-Feijo/leago/internal"
)

// RegionClient groups the League of Legends APIs that are scoped to a routing region (e.g. EUROPE, AMERICAS).
type RegionClient struct {
	Match      *match.RegionClient
	Tournament *tournament.RegionClient
}

func NewRegionClient(baseClient *internal.Client) *RegionClient {
	c := &RegionClient{
		Match:      match.NewRegionClient(baseClient),
		Tournament: tournament.NewRegionClient(baseClient),
	}
	return c
}
