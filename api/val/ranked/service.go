package ranked

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetLeaderboard = "ValRanked.GetLeaderboard"
)

// GetLeaderboard returns the leaderboard for the competitive queue.
func (pc *RegionClient) GetLeaderboard(
	ctx context.Context,
	actID string,
	endpointOpts []RankedOption,
	opts ...options.PublicOption,
) (Leaderboard, error) {

	endpoint := fmt.Sprintf(
		"/val/ranked/v1/leaderboards/by-act/%s",
		url.PathEscape(actID),
	)

	defaultOpts := append(
		[]internal.RequestOption{internal.WithAPIMethod(MethodGetLeaderboard)},
		rankedOptionsToRequestOptions(endpointOpts)...,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Leaderboard](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
