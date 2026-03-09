package clash

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"testing"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/internal/mock"
	"github.com/Gustavo-Feijo/leago/regions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedPlayer = []Player{
		{
			Puuid:    "test-puuid",
			TeamID:   "456",
			Position: "TOP",
			Role:     "CAPTAIN",
		},
	}

	playerJSON = `[{
		"puuid":"test-puuid",
		"teamId":"456",
		"position":"TOP",
		"role":"CAPTAIN"
	}]`

	expectedTeam = Team{
		ID:           "test-team-id",
		TournamentID: 1,
		Name:         "Test Team",
		IconID:       10,
		Tier:         2,
		Captain:      "captain-puuid",
		Abbreviation: "TT",
		Players: []TeamPlayer{
			{
				Puuid:    "player-puuid",
				Position: "TOP",
				Role:     "MEMBER",
			},
		},
	}

	teamJSON = `{
		"id":"test-team-id",
		"tournamentId":1,
		"name":"Test Team",
		"iconId":10,
		"tier":2,
		"captain":"captain-puuid",
		"abbreviation":"TT",
		"players":[{
			"puuid":"player-puuid",
			"position":"TOP",
			"role":"MEMBER"
		}]
	}`

	expectedTournament = Tournament{
		ID:               1,
		ThemeID:          2,
		NameKey:          "clash",
		NameKeySecondary: "secondary",
		Schedule: []TournamentPhase{
			{
				ID:               10,
				RegistrationTime: 1000,
				StartTime:        2000,
				Cancelled:        false,
			},
		},
	}

	tournamentJSON = `{
		"id":1,
		"themeId":2,
		"nameKey":"clash",
		"nameKeySecondary":"secondary",
		"schedule":[{
			"id":10,
			"registrationTime":1000,
			"startTime":2000,
			"cancelled":false
		}]
	}`

	expectedTournaments = []Tournament{expectedTournament}

	tournamentsJSON = fmt.Sprintf("[%s]", tournamentJSON)
)

func newTestPlatformClient(statusCode int, responseBody string) (*PlatformClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
	return NewPlatformClient(baseClient), mockDoer
}

func TestGetPlayerByPUUID(t *testing.T) {
	tests := []struct {
		name string

		puuid string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult []Player

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			puuid: "test-puuid",

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "unmatched json",

			puuid: "test-puuid",

			statusCode:   http.StatusOK,
			responseBody: `{"puuid":"shouldbearray"}`,

			wantErr:     true,
			wantRiotErr: false,
		},
		{
			name: "success",

			puuid: "test-puuid",

			statusCode:   http.StatusOK,
			responseBody: playerJSON,

			expectedPath: "/lol/clash/v1/players/by-puuid/test-puuid",

			expectedResult: expectedPlayer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc, mockDoer := newTestPlatformClient(tt.statusCode, tt.responseBody)
			resp, err := pc.GetPlayerByPUUID(context.Background(), tt.puuid)

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
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetTeamByID(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, teamJSON)
	resp, err := pc.GetTeamByID(context.Background(), "test-team-id")

	require.NoError(t, err)

	assert.Equal(t, "/lol/clash/v1/teams/test-team-id", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedTeam, resp)
}

func TestGetTournaments(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, tournamentsJSON)
	resp, err := pc.GetTournaments(context.Background())

	require.NoError(t, err)

	assert.Equal(t, "/lol/clash/v1/tournaments", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedTournaments, resp)
}

func TestGetTournamentByTeamID(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, tournamentJSON)
	resp, err := pc.GetTournamentByTeamID(context.Background(), "test-team-id")

	require.NoError(t, err)

	assert.Equal(t, "/lol/clash/v1/tournaments/by-team/test-team-id", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedTournament, resp)
}

func TestGetTournamentByID(t *testing.T) {
	pc, mockDoer := newTestPlatformClient(http.StatusOK, tournamentJSON)
	resp, err := pc.GetTournamentByID(context.Background(), "123")

	require.NoError(t, err)

	assert.Equal(t, "/lol/clash/v1/tournaments/123", mockDoer.CapturedReq.URL.Path)
	assert.Equal(t, expectedTournament, resp)
}
