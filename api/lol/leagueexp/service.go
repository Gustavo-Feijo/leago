package leagueexp

import (
	"context"
	"fmt"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetLeague = "LolLeagueexp.GetLeague"
)

// GetLeague returns all league entries based on the params.
// More consistent than league, which has some weird separations (Separate APIs for upper divisions).
func (pc *PlatformClient) GetLeague(
	ctx context.Context,
	queue Queue, tier Tier,
	division Division,
	endpointOpts []GetLeagueOption,
	opts ...options.PublicOption,
) ([]Entry, error) {
	endpoint := fmt.Sprintf(
		"/lol/league-exp/v4/entries/%s/%s/%s",
		queue,
		tier,
		division,
	)

	defaultOpts := append(
		[]internal.RequestOption{internal.WithAPIMethod(MethodGetLeague)},
		getLeagueOptionsToRequestOptions(endpointOpts)...,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[[]Entry](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
