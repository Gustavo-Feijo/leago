package spectator

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetGameByPUUID = "Spectator.GetGameByPUUID"
)

// GetGameByPUUID gets the game info for a given PUUID.
func (pc *PlatformClient) GetGameByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) (CurrentGameInfo, error) {
	endpoint := fmt.Sprintf(
		"/lol/spectator/v5/active-games/by-summoner/%s",
		url.PathEscape(puuid),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetGameByPUUID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[CurrentGameInfo](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
