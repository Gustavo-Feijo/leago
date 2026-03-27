package champion

import "github.com/Gustavo-Feijo/leago/internal"

type RegionClient struct {
	client   *internal.Client
	version  string
	language string
}

func NewRegionClient(base *internal.Client, version, language string) *RegionClient {
	return &RegionClient{
		client:   base,
		version:  version,
		language: language,
	}
}
