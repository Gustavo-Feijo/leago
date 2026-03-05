package summoner

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetSummonerByPUUID = "Summoner.GetSummonerByPUUID"
)

// GetSummonerByPUUID retrieves the League Of Legends summoner by the PUUID.
func (pc *PlatformClient) GetSummonerByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) (Summoner, error) {
	endpoint := fmt.Sprintf(
		"/lol/summoner/v4/summoners/by-puuid/%s",
		url.PathEscape(puuid),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetSummonerByPUUID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Summoner](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
