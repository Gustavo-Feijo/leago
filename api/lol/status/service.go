package status

import (
	"context"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetStatus = "LolStatus.GetStatus"
)

// GetStatus retrieves the League Of Legends status for the platforms.
func (pc *PlatformClient) GetStatus(
	ctx context.Context,
	opts ...options.PublicOption,
) (ServiceStatus, error) {
	endpoint := "/lol/status/v4/platform-data"

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetStatus),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[ServiceStatus](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
