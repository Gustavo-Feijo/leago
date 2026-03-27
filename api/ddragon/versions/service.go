package versions

import (
	"context"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetVersion = "DDVersions.GetVersions"
)

// GetVersions returns a list of all available versions.
func (rc *RegionClient) GetVersions(
	ctx context.Context,
	opts ...options.PublicOption,
) ([]string, error) {
	endpoint := "/api/versions.json"

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetVersion),
	}

	uri := rc.client.GetDDragonURL(endpoint)
	return internal.Request[[]string](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
