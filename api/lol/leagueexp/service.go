package leagueexp

import (
	"context"
	"fmt"
	"leago/internal"
	"leago/options"
)

const (
	MethodGetLeague = "Leagueexp.GetLeague"
)

// GetLeague returns all league entries based on the params.
// More consistent than league, which has some weird separations (Separate APIs for upper divisions).
func (pc *PlatformClient) GetLeague(ctx context.Context, queue Queue, tier Tier, division Division, opts ...options.PublicOption) (LeagueResponse, error) {
	endpoint := fmt.Sprintf(
		"/lol/league-exp/v4/entries/%s/%s/%s",
		queue,
		tier,
		division,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[LeagueResponse](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(
			[]internal.RequestOption{
				internal.WithApiMethod(MethodGetLeague),
			},
			opts,
		)...,
	)
}
