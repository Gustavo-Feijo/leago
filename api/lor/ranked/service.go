package ranked

import (
	"context"
	"leago/internal"
	"leago/options"
)

const (
	MethodGetLeaderboards = "LorRanked.GetLeaderboards"
)

// GetLeaderboards gets the LOR leaderboard.
func (rc *RegionClient) GetLeaderboards(
	ctx context.Context,
	opts ...options.PublicOption,
) (Leaderboard, error) {
	endpoint := "/lor/ranked/v1/leaderboards"

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetLeaderboards),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[Leaderboard](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
