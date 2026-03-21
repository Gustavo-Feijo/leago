package match

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetMatchesByPUUID       = "ValMatches.GetMatchesByPUUID"
	MethodGetMatchByID            = "ValMatches.GetMatchByID"
	MethodGetRecentMatchesByQueue = "ValMatches.GetRecentMatchesByQueue"
)

// GetMatchesByPUUID returns a list of match IDs for a given PUUID.
func (rc *RegionClient) GetMatchesByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) (MatchList, error) {
	endpoint := fmt.Sprintf(
		"/val/match/v1/matchlists/by-puuid/%s",
		url.PathEscape(puuid),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetMatchesByPUUID),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[MatchList](
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
		"/val/match/v1/matches/%s",
		url.PathEscape(matchID),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetMatchByID),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[Match](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetRecentMatchesByQueue returns the list of matches that have been completed recently.
func (rc *RegionClient) GetRecentMatchesByQueue(
	ctx context.Context,
	queue Queue,
	opts ...options.PublicOption,
) (RecentMatches, error) {
	endpoint := fmt.Sprintf(
		"/val/match/v1/recent-matches/by-queue/%s",
		queue,
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetRecentMatchesByQueue),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[RecentMatches](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
