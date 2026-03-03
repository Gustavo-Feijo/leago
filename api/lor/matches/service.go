package matches

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetMatchesByPUUID = "LorMatches.GetMatchesByPUUID"
	MethodGetMatchByID      = "LorMatches.GetMatchByID"
)

// GetMatchesByPUUID returns a list of match IDs for a given PUUID.
func (rc *RegionClient) GetMatchesByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) ([]string, error) {
	endpoint := fmt.Sprintf(
		"/lor/match/v1/matches/by-puuid/%s/ids",
		url.PathEscape(puuid),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetMatchesByPUUID),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[[]string](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetMatchByID returns a given match data.
func (rc *RegionClient) GetMatchByID(
	ctx context.Context,
	matchID string,
	opts ...options.PublicOption,
) (Match, error) {
	endpoint := fmt.Sprintf(
		"/lor/match/v1/matches/%s",
		url.PathEscape(matchID),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetMatchByID),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[Match](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
