package clash

import (
	"context"
	"fmt"
	"leago/internal"
	"net/url"
)

// GetByPUUID returns the player champion mastery information got by their PUUID.
func (pc *PlatformClient) GetPlayerByPUUID(ctx context.Context, puuid string) (PlayersResponse, error) {
	endpoint := fmt.Sprintf(
		"/lol/clash/v1/players/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[PlayersResponse](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod("Clash.GetPlayerByPUUID"),
	)
}

// GetTeamByID returns the team got by their ID.
func (pc *PlatformClient) GetTeamByID(ctx context.Context, teamId string) (Team, error) {
	endpoint := fmt.Sprintf(
		"/lol/clash/v1/teams/%s",
		url.PathEscape(teamId),
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Team](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod("Clash.GetTeamByID"),
	)
}

// GetTournaments returns a list of active and upcoming tournaments.
func (pc *PlatformClient) GetTournaments(ctx context.Context) (TournamentsResponse, error) {
	endpoint := "/lol/clash/v1/tournaments"

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[TournamentsResponse](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod("Clash.GetTournaments"),
	)
}

// GetTournamentByTeamID returns a tournament got by the teamId.
func (pc *PlatformClient) GetTournamentByTeamID(ctx context.Context, teamId string) (Tournament, error) {
	endpoint := fmt.Sprintf(
		"/lol/clash/v1/tournaments/by-team/%s",
		teamId,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Tournament](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod("Clash.GetTournamentByTeamID"),
	)
}

// GetTournamentByID returns a tournament got by the tournamentId.
func (pc *PlatformClient) GetTournamentByID(ctx context.Context, tournamentId string) (Tournament, error) {
	endpoint := fmt.Sprintf(
		"/lol/clash/v1/tournaments/%s",
		tournamentId,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Tournament](
		ctx,
		pc.client,
		uri,
		internal.WithApiMethod("Clash.GetTournamentByID"),
	)
}
