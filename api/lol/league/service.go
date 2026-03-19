package league

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetMasterLeague      = "League.GetMasterLeague"
	MethodGetGrandmasterLeague = "League.GetGrandmasterLeague"
	MethodGetChallengerLeague  = "League.GetChallengerLeague"

	MethodGetLeagueEntries        = "League.GetLeagueEntries"
	MethodGetLeagueEntriesByPUUID = "League.GetLeagueEntriesByPUUID"
	MethodGetLeagueByID           = "League.GetLeagueByID"
)

// GetMasterLeague returns all entries on the master league.
func (pc *PlatformClient) GetMasterLeague(
	ctx context.Context,
	queue Queue,
	opts ...options.PublicOption,
) (RawLeague, error) {
	endpoint := fmt.Sprintf(
		"/lol/league/v4/masterleagues/by-queue/%s",
		queue,
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetMasterLeague),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[RawLeague](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetGrandmasterLeague returns all entries on the grandmaster league.
func (pc *PlatformClient) GetGrandmasterLeague(
	ctx context.Context,
	queue Queue,
	opts ...options.PublicOption,
) (RawLeague, error) {
	endpoint := fmt.Sprintf(
		"/lol/league/v4/grandmasterleagues/by-queue/%s",
		queue,
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetGrandmasterLeague),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[RawLeague](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetChallengerLeague returns all entries on the challenger league.
func (pc *PlatformClient) GetChallengerLeague(
	ctx context.Context,
	queue Queue,
	opts ...options.PublicOption,
) (RawLeague, error) {
	endpoint := fmt.Sprintf(
		"/lol/league/v4/challengerleagues/by-queue/%s",
		queue,
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetChallengerLeague),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[RawLeague](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetLeagueEntries returns all entries for a given league up to Diamond I.
func (pc *PlatformClient) GetLeagueEntries(
	ctx context.Context,
	queue Queue,
	tier Tier,
	division Division,
	endpointOpts []GetLeagueOption,
	opts ...options.PublicOption,
) ([]Entry, error) {

	endpoint := fmt.Sprintf(
		"/lol/league/v4/entries/%s/%s/%s",
		queue,
		tier,
		division,
	)

	defaultOpts := append(
		[]internal.RequestOption{internal.WithAPIMethod(MethodGetLeagueEntries)},
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

// GetLeagueEntriesByPUUID returns the league entries for a given player got by the PUUID.
func (pc *PlatformClient) GetLeagueEntriesByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) ([]Entry, error) {
	endpoint := fmt.Sprintf(
		"/lol/league/v4/entries/by-puuid/%s",
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

// GetLeagueByID returns a given league entries got by its leagueID.
func (pc *PlatformClient) GetLeagueByID(
	ctx context.Context,
	leagueID string,
	opts ...options.PublicOption,
) (RawLeague, error) {
	endpoint := fmt.Sprintf(
		"/lol/league/v4/leagues/%s",
		leagueID,
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetLeagueByID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[RawLeague](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
