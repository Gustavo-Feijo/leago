package tournament

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodCreateCodes                    = "LolTournament.CreateCodes"
	MethodGetCodes                       = "LolTournament.GetCodes"
	MethodUpdateCodes                    = "LolTournament.UpdateCodes"
	MethodGetGamesByCode                 = "LolTournament.GetGamesByCode"
	MethodGetLobbyEventsByTournamentCode = "LolTournament.GetLobbyEventsByTournamentCode"
	MethodCreateProvider                 = "LolTournament.CreateProvider"
	MethodCreateTournament               = "LolTournament.CreateTournament"
)

// CreateCodes creates tournament codes.
func (rc *RegionClient) CreateCodes(
	ctx context.Context,
	tournamentID int64,
	body *TournamentCodePayload,
	stub bool,
	endpointOpts []CreateCodeOption,
	opts ...options.PublicOption,
) ([]string, error) {
	endpoint := "/lol/tournament/v5/codes"
	if stub {
		endpoint = "/lol/tournament-stub/v5/codes"
	}

	defaultOpts := append(
		[]internal.RequestOption{
			internal.WithAPIMethod(MethodCreateCodes),
			internal.WithHTTPMethod(http.MethodPost),
			internal.WithBody(body),
			internal.WithParam("tournamentId", strconv.FormatInt(tournamentID, 10)),
		},
		createCodeOptionsToRequestOptions(endpointOpts)...,
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[[]string](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetCodes return the tournament dto for a given tournament code.
func (rc *RegionClient) GetCodes(
	ctx context.Context,
	tournamentCode string,
	stub bool,
	opts ...options.PublicOption,
) (TournamentCode, error) {
	baseEndpoint := "/lol/tournament/v5/codes/%s"
	if stub {
		baseEndpoint = "/lol/tournament-stub/v5/codes/%s"
	}
	endpoint := fmt.Sprintf(baseEndpoint, url.PathEscape(tournamentCode))

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetCodes),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[TournamentCode](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// UpdateCodes updates some fields for a code.
func (rc *RegionClient) UpdateCodes(
	ctx context.Context,
	tournamentCode string,
	body *PutTournamentCodePayload,
	opts ...options.PublicOption,
) (string, error) {
	endpoint := fmt.Sprintf(
		"/lol/tournament/v5/codes/%s",
		url.PathEscape(tournamentCode),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodUpdateCodes),
		internal.WithHTTPMethod(http.MethodPut),
		internal.WithBody(body),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[string](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetGamesByCode return game details for a given tournament.
func (rc *RegionClient) GetGamesByCode(
	ctx context.Context,
	tournamentCode string,
	opts ...options.PublicOption,
) (TournamentGame, error) {
	endpoint := fmt.Sprintf(
		"/lol/tournament/v5/games/by-code/%s",
		url.PathEscape(tournamentCode),
	)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodUpdateCodes),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[TournamentGame](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetLobbyEventsByTournamentCode return the lobby events for a given tournament.
func (rc *RegionClient) GetLobbyEventsByTournamentCode(
	ctx context.Context,
	tournamentCode string,
	stub bool,
	opts ...options.PublicOption,
) (LobbyEventWrapper, error) {
	baseEndpoint := "/lol/tournament/v5/lobby-events/by-code/%s"
	if stub {
		baseEndpoint = "/lol/tournament-stub/v5/lobby-events/by-code/%s"
	}
	endpoint := fmt.Sprintf(baseEndpoint, url.PathEscape(tournamentCode))

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetLobbyEventsByTournamentCode),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[LobbyEventWrapper](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// CreateProvider creates a tournament provider and returns its ID.
func (rc *RegionClient) CreateProvider(
	ctx context.Context,
	body *ProviderPayload,
	stub bool,
	opts ...options.PublicOption,
) (int, error) {
	endpoint := "/lol/tournament/v5/providers"
	if stub {
		endpoint = "/lol/tournament-stub/v5/providers"
	}

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodCreateProvider),
		internal.WithHTTPMethod(http.MethodPost),
		internal.WithBody(body),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[int](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// CreateTournament creates a tournament and returns its ID.
func (rc *RegionClient) CreateTournament(
	ctx context.Context,
	body *TournamentPayload,
	stub bool,
	opts ...options.PublicOption,
) (int, error) {
	endpoint := "/lol/tournament/v5/providers"
	if stub {
		endpoint = "/lol/tournament-stub/v5/providers"
	}

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodCreateTournament),
		internal.WithHTTPMethod(http.MethodPost),
		internal.WithBody(body),
	}

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[int](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
