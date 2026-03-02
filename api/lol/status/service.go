package status

import (
	"context"
	"leago/internal"
	"leago/options"
)

const (
	MethodGetStatus = "Status.GetStatus"
)

// GetStatus retrieves the League Of Legends status for the platforms.
func (pc *PlatformClient) GetStatus(
	ctx context.Context,
	opts ...options.PublicOption,
) (ServiceStatus, error) {
	endpoint := "/lol/status/v4/platform-data"

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetStatus),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[ServiceStatus](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
