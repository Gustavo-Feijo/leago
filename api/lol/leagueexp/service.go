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
func (pc *PlatformClient) GetLeague(
	ctx context.Context,
	queue Queue, tier Tier,
	division Division,
	endpointOpts []GetLeagueOption,
	opts ...options.PublicOption,
) (LeagueResponse, error) {
	endpoint := fmt.Sprintf(
		"/lol/league-exp/v4/entries/%s/%s/%s",
		queue,
		tier,
		division,
	)

	defaultOpts := append(
		[]internal.RequestOption{internal.WithApiMethod(MethodGetLeague)},
		getLeagueOptionsToRequestOptions(endpointOpts)...,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[LeagueResponse](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
