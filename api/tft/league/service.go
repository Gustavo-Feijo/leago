package league

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetMasterLeague      = "TFTLeague.GetMasterLeague"
	MethodGetGrandmasterLeague = "TFTLeague.GetGrandmasterLeague"
	MethodGetChallengerLeague  = "TFTLeague.GetChallengerLeague"

	MethodGetLeagueByID           = "TFTLeague.GetLeagueByID"
	MethodGetLeagueEntries        = "TFTLeague.GetLeagueEntries"
	MethodGetLeagueEntriesByPUUID = "TFTLeague.GetLeagueEntriesByPUUID"
	MethodGetRatedLadder          = "TFTLeague.GetRatedLadder"
)

// GetMasterLeague returns the master league.
func (pc *PlatformClient) GetMasterLeague(
	ctx context.Context,
	endpointOpts []UpperLeagueOption,
	opts ...options.PublicOption,
) (List, error) {
	endpoint := "/tft/league/v1/master"

	defaultOpts := append(
		[]internal.RequestOption{internal.WithAPIMethod(MethodGetMasterLeague)},
		upperLeagueOptionsToRequestOptions(endpointOpts)...,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[List](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetGrandmasterLeague returns the grandmaster league.
func (pc *PlatformClient) GetGrandmasterLeague(
	ctx context.Context,
	endpointOpts []UpperLeagueOption,
	opts ...options.PublicOption,
) (List, error) {
	endpoint := "/tft/league/v1/grandmaster"

	defaultOpts := append(
		[]internal.RequestOption{internal.WithAPIMethod(MethodGetGrandmasterLeague)},
		upperLeagueOptionsToRequestOptions(endpointOpts)...,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[List](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetChallengerLeague returns the challenger league.
func (pc *PlatformClient) GetChallengerLeague(
	ctx context.Context,
	endpointOpts []UpperLeagueOption,
	opts ...options.PublicOption,
) (List, error) {
	endpoint := "/tft/league/v1/challenger"

	defaultOpts := append(
		[]internal.RequestOption{internal.WithAPIMethod(MethodGetChallengerLeague)},
		upperLeagueOptionsToRequestOptions(endpointOpts)...,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[List](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetLeagueByID returns a league by the ID.
func (pc *PlatformClient) GetLeagueByID(
	ctx context.Context,
	leagueID string,
	opts ...options.PublicOption,
) (List, error) {

	endpoint := fmt.Sprintf(
		"/tft/league/v1/leagues/%s",
		url.PathEscape(leagueID),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetLeagueByID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[List](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetLeagueEntries returns all entries for a given league up to Diamond I.
func (pc *PlatformClient) GetLeagueEntries(
	ctx context.Context,
	tier Tier,
	division Division,
	endpointOpts []LeagueOption,
	opts ...options.PublicOption,
) ([]Entry, error) {

	endpoint := fmt.Sprintf(
		"/tft/league/v1/entries/%s/%s",
		tier,
		division,
	)

	defaultOpts := append(
		[]internal.RequestOption{internal.WithAPIMethod(MethodGetLeagueEntries)},
		leagueOptionsToRequestOptions(endpointOpts)...,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[[]Entry](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetLeagueEntriesByPUUID returns all league entries based on the puuid.
func (pc *PlatformClient) GetLeagueEntriesByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) ([]Entry, error) {
	endpoint := fmt.Sprintf(
		"/tft/league/v1/by-puuid/%s",
		url.PathEscape(puuid),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetLeagueEntriesByPUUID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[[]Entry](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetRatedLadder returns the top ladder for a given queue.
func (pc *PlatformClient) GetRatedLadder(
	ctx context.Context,
	queue LadderQueue,
	opts ...options.PublicOption,
) ([]RatedLadderEntry, error) {

	endpoint := fmt.Sprintf(
		"/tft/league/v1/rated-ladders/%s/top",
		queue,
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetRatedLadder),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[[]RatedLadderEntry](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
