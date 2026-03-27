package summoner

import (
	"context"
	"fmt"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetSummonerSpells = "DDSummoner.GetSummonerSpells"
)

// GetSummonerSpells returns a list of all summonerspells.
func (rc *RegionClient) GetSummonerSpells(
	ctx context.Context,
	opts ...options.PublicOption,
) (SummonerResponse, error) {
	endpoint := fmt.Sprintf("/cdn/%s/data/%s/summoner.json", rc.version, rc.language)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetSummonerSpells),
	}

	uri := rc.client.GetDDragonURL(endpoint)
	return internal.Request[SummonerResponse](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
