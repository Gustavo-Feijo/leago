package matches

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetMatchesByPUUID    = "LolMatches.GetMatchesByPUUID"
	MethodGetReplaysByPUUID    = "LolMatches.GetReplaysByPUUID"
	MethodGetMatchByID         = "LolMatches.GetMatchByID"
	MethodGetMatchTimelineByID = "LolMatches.GetMatchTimelineByID"
)

// GetMatchesByPUUID returns a list of match IDs for a given PUUID.
func (rc *RegionClient) GetMatchesByPUUID(
	ctx context.Context,
	puuid string,
	endpointOpts []GetMatchesByPUUIDOption,
	opts ...options.PublicOption,
) ([]string, error) {
	endpoint := fmt.Sprintf(
		"/lol/match/v5/matches/by-puuid/%s/ids",
		url.PathEscape(puuid),
	)

	defaultOpts := append(
		[]internal.RequestOption{internal.WithApiMethod(MethodGetMatchesByPUUID)},
		getMatchByPUUIDOptionsToRequestOptions(endpointOpts)...,
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[[]string](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetReplaysByPUUID returns the URL of the replay files from a given player.
func (rc *RegionClient) GetReplaysByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) (Replays, error) {
	endpoint := fmt.Sprintf(
		"/lol/match/v5/matches/by-puuid/%s/replays",
		url.PathEscape(puuid),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetReplaysByPUUID),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[Replays](
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
		"/lol/match/v5/matches/%s",
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

// GetMatchByID returns a given match data.
func (rc *RegionClient) GetMatchTimelineByID(
	ctx context.Context,
	matchID string,
	opts ...options.PublicOption,
) (Timeline, error) {
	endpoint := fmt.Sprintf(
		"/lol/match/v5/matches/%s/timeline",
		url.PathEscape(matchID),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetMatchTimelineByID),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[Timeline](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
