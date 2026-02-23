package champion

import (
	"context"
	"leago/internal"
)

// GetRotation returns the current free champion rotation.
func (pc *PlatformClient) GetRotation(ctx context.Context) (Rotation, error) {
	endpoint := "/lol/platform/v3/champion-rotations"

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Rotation](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod("Champion.GetRotation"),
	)
}
