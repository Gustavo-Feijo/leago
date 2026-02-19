package account

import "leago/internal"

type RegionClient struct {
	client *internal.Client
}

func NewRegionClient(base *internal.Client) *RegionClient {
	return &RegionClient{
		base,
	}
}
