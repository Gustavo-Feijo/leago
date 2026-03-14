package tft

import (
	"github.com/Gustavo-Feijo/leago/api/tft/match"
	"github.com/Gustavo-Feijo/leago/internal"
)

type RegionClient struct {
	Match *match.RegionClient
}

func NewRegionClient(baseClient *internal.Client) *RegionClient {
	c := &RegionClient{
		Match: match.NewRegionClient(baseClient),
	}
	return c
}
