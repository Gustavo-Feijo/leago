package clash

import (
	"context"
	"fmt"
	"leago/internal"
	"leago/options"
	"net/url"
)

const (
	MethodGetPlayerByPUUID      = "Clash.GetPlayerByPUUID"
	MethodGetTeamByID           = "Clash.GetTeamByID"
	MethodGetTournaments        = "Clash.GetTournaments"
	MethodGetTournamentByTeamID = "Clash.GetTournamentByTeamID"
	MethodGetTournamentByID     = "Clash.GetTournamentByID"
)

// GetByPUUID returns the player champion mastery information got by their PUUID.
func (pc *PlatformClient) GetPlayerByPUUID(
	ctx context.Context,
	puuid string,
	opts ...options.PublicOption,
) (PlayersResponse, error) {
	endpoint := fmt.Sprintf(
		"/lol/clash/v1/players/by-puuid/%s",
		url.PathEscape(puuid),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetPlayerByPUUID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[PlayersResponse](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetTeamByID returns the team got by their ID.
func (pc *PlatformClient) GetTeamByID(
	ctx context.Context,
	teamID string,
	opts ...options.PublicOption,
) (Team, error) {
	endpoint := fmt.Sprintf(
		"/lol/clash/v1/teams/%s",
		url.PathEscape(teamID),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetTeamByID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Team](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetTournaments returns a list of active and upcoming tournaments.
func (pc *PlatformClient) GetTournaments(
	ctx context.Context,
	opts ...options.PublicOption,
) (TournamentsResponse, error) {
	endpoint := "/lol/clash/v1/tournaments"

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetTournaments),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[TournamentsResponse](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetTournamentByTeamID returns a tournament got by the teamID.
func (pc *PlatformClient) GetTournamentByTeamID(
	ctx context.Context,
	teamID string,
	opts ...options.PublicOption,
) (Tournament, error) {
	endpoint := fmt.Sprintf(
		"/lol/clash/v1/tournaments/by-team/%s",
		url.PathEscape(teamID),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetTournamentByTeamID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Tournament](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetTournamentByID returns a tournament got by the tournamentID.
func (pc *PlatformClient) GetTournamentByID(
	ctx context.Context,
	tournamentID string,
	opts ...options.PublicOption,
) (Tournament, error) {
	endpoint := fmt.Sprintf(
		"/lol/clash/v1/tournaments/%s",
		url.PathEscape(tournamentID),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithApiMethod(MethodGetTournamentByID),
	}

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Tournament](
		ctx,
		pc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
