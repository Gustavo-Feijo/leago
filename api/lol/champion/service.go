package champion

import (
	"context"
	"leago/internal"
	"leago/options"
)

const (
	MethodGetRotation = "Champion.GetRotation"
)

// GetRotation returns the current free champion rotation.
func (pc *PlatformClient) GetRotation(ctx context.Context, opts ...options.PublicOption) (Rotation, error) {
	endpoint := "/lol/platform/v3/champion-rotations"

	uri := pc.client.GetURL(endpoint)

	return internal.AuthRequest[Rotation](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(
			[]internal.RequestOption{internal.WithApiMethod(MethodGetRotation)},
			opts,
		)...,
	)
}
