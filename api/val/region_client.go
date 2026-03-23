package val

import (
	"github.com/Gustavo-Feijo/leago/api/val/content"
	"github.com/Gustavo-Feijo/leago/api/val/match"
	"github.com/Gustavo-Feijo/leago/api/val/ranked"
	"github.com/Gustavo-Feijo/leago/api/val/status"
	"github.com/Gustavo-Feijo/leago/internal"
)

// RegionClient groups the Valorant APIs that are scoped to a routing region (e.g. br, eu).
type RegionClient struct {
	Content *content.RegionClient
	Match   *match.RegionClient
	Ranked  *ranked.RegionClient
	Status  *status.RegionClient
}

func NewRegionClient(baseClient *internal.Client) *RegionClient {
	c := &RegionClient{
		Content: content.NewRegionClient(baseClient),
		Match:   match.NewRegionClient(baseClient),
		Ranked:  ranked.NewRegionClient(baseClient),
		Status:  status.NewRegionClient(baseClient),
	}
	return c
}
