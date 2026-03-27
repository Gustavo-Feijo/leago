package tournament

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedCodes = []string{
		"tournament1",
		"tournament2",
	}

	codesJSON = `
		[
		"tournament1",
		"tournament2"
		]
	`

	expectedCode = TournamentCode{
		Code:         "code1",
		LobbyName:    "lobby",
		MetaData:     "meta",
		Password:     "pass",
		TeamSize:     5,
		ProviderID:   1,
		PickType:     PickBlind,
		TournamentID: 10,
		ID:           100,
		Region:       RegionBR,
		Map:          MapSummonersRift,
		Participants: []string{"p1", "p2"},
	}

	codeJSON = `{
		"code":"code1",
		"lobbyName":"lobby",
		"metaData":"meta",
		"password":"pass",
		"teamSize":5,
		"providerId":1,
		"pickType":"BLIND_PICK",
		"tournamentId":10,
		"id":100,
		"region":"BR",
		"map":"SUMMONERS_RIFT",
		"participants":["p1","p2"]
	}`

	expectedGame = TournamentGame{
		StartTime: 123456,
		WinningTeam: []TournamentTeam{
			{PUUID: "w1"},
		},
		LosingTeam: []TournamentTeam{
			{PUUID: "l1"},
		},
		ShortCode: "short",
		MetaData:  "meta",
		GameID:    999,
		GameName:  "game",
		GameType:  "type",
		GameMap:   "map",
		GameMode:  "mode",
		Region:    RegionBR,
	}

	gameJSON = `{
		"startTime":123456,
		"winningTeam":[{"puuid":"w1"}],
		"losingTeam":[{"puuid":"l1"}],
		"shortCode":"short",
		"metaData":"meta",
		"gameId":999,
		"gameName":"game",
		"gameType":"type",
		"gameMap":"map",
		"gameMode":"mode",
		"region":"BR"
	}`

	expectedEvents = LobbyEventWrapper{
		EventList: []LobbyEvent{
			{
				Timestamp: internal.LobbyTime{Time: time.Date(2026, 03, 22, 17, 0, 0, 0, time.UTC)},
				EventType: "JOIN",
			},
		},
	}

	eventsJSON = `{
		"eventList":[
			{
				"timestamp":"Sun Mar 22 17:00:00 UTC 2026",
				"eventType":"JOIN"
			}
		]
	}`

	expectedProvider = 123

	providerJSON = `123`

	expectedTournament = 456

	tournamentJSON = `456`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.RegionAmericas), "apiKey", internal.WithHTTP(mockDoer))

	return NewRegionClient(baseClient), mockDoer
}

func TestCreateCodes(t *testing.T) {
	tests := []struct {
		name string

		tournamentID int64
		stub         bool
		opts         []CreateCodeOption

		statusCode   int
		responseBody string

		expectedPath  string
		expectedQuery map[string]string

		expectedResult []string

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			tournamentID: 1,

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "invalid json",

			tournamentID: 2,

			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"shouldbevalid"}`,

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "success",

			tournamentID: 3,
			stub:         true,
			opts:         []CreateCodeOption{WithCount(5)},

			statusCode:   http.StatusOK,
			responseBody: codesJSON,

			expectedPath: "/lol/tournament-stub/v5/codes",
			expectedQuery: map[string]string{
				"count": "5",
			},

			expectedResult: expectedCodes,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)
			resp, err := rc.CreateCodes(
				context.Background(),
				tt.tournamentID,
				&TournamentCodePayload{
					SpectatorType: SpectatorAll,
					TeamSize:      5,
					PickType:      PickBlind,
					MapType:       MapSummonersRift,
					EnoughPlayers: true,
				},
				tt.stub,
				tt.opts,
			)

			if tt.wantErr {
				require.Error(t, err)

				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}

				return
			}

			require.NoError(t, err)

			assert.Equal(t, tt.expectedPath, mockDoer.CapturedReq.URL.Path)

			query := mockDoer.CapturedReq.URL.Query()
			for k, v := range tt.expectedQuery {
				assert.Equal(t, v, query.Get(k))
			}

			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetCodes(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, codeJSON)

	resp, err := rc.GetCodes(context.Background(), "code1", true)

	require.NoError(t, err)

	assert.Equal(t, "/lol/tournament-stub/v5/codes/code1", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedCode, resp)
}

func TestUpdateCodes(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, `"updated"`)

	resp, err := rc.UpdateCodes(
		context.Background(),
		"code1",
		&PutTournamentCodePayload{
			SpectatorType: SpectatorAll,
			PickType:      PickBlind,
			MapType:       MapSummonersRift,
		},
	)

	require.NoError(t, err)

	assert.Equal(t, "/lol/tournament/v5/codes/code1", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, "updated", resp)
}

func TestGetGamesByCode(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, gameJSON)

	resp, err := rc.GetGamesByCode(context.Background(), "code1")

	require.NoError(t, err)

	assert.Equal(t, "/lol/tournament/v5/games/by-code/code1", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedGame, resp)
}

func TestGetLobbyEventsByTournamentCode(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, eventsJSON)

	resp, err := rc.GetLobbyEventsByTournamentCode(context.Background(), "code1", true)

	require.NoError(t, err)

	assert.Equal(t, "/lol/tournament-stub/v5/lobby-events/by-code/code1", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedEvents, resp)
}

func TestCreateProvider(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, providerJSON)

	resp, err := rc.CreateProvider(
		context.Background(),
		&ProviderPayload{
			Region: RegionBR,
			URL:    "http://callback",
		},
		true,
	)

	require.NoError(t, err)

	assert.Equal(t, "/lol/tournament-stub/v5/providers", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedProvider, resp)
}

func TestCreateTournament(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, tournamentJSON)

	resp, err := rc.CreateTournament(
		context.Background(),
		&TournamentPayload{
			Name:     "tournament",
			Provider: 1,
		},
		true,
	)

	require.NoError(t, err)

	assert.Equal(t, "/lol/tournament-stub/v5/tournaments", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedTournament, resp)
}
