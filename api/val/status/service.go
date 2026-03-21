package status

import (
	"context"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetStatus = "ValStatus.GetStatus"
)

// GetStatus retrieves the Valorant status for the platforms.
func (rc *RegionClient) GetStatus(
	ctx context.Context,
	opts ...options.PublicOption,
) (ServiceStatus, error) {
	endpoint := "/val/status/v1/platform-data"

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetStatus),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[ServiceStatus](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
