package summoner

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetSummonerByPUUID = "TFTSummoner.GetSummonerByPUUID"
)

// GetSummonerByPUUID retrieves the TFT summoner by the PUUID.
func (pc *PlatformClient) GetSummonerByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) (Summoner, error) {
	endpoint := fmt.Sprintf(
		"/tft/summoner/v1/summoners/by-puuid/%s",
		url.PathEscape(puuid),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetSummonerByPUUID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Summoner](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
