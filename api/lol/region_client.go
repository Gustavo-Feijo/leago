package lol

import (
	"github.com/Gustavo-Feijo/leago/api/lol/match"
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
