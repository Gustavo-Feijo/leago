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

func TestGetPlayerByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		httpErr        error
		responseBody   string
		expectedResult []Player
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			puuid:        "test-puuid",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			puuid:        "test-puuid",
			statusCode:   http.StatusOK,
			responseBody: `{"puuid":"shouldbearray"}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:       "success",
			puuid:      "test-puuid",
			statusCode: http.StatusOK,
			responseBody: `[{
				"puuid":"test-puuid",
				"teamId":"456",
				"position":"TOP",
				"role":"CAPTAIN"
			}]`,
			expectedResult: []Player{
				{
					Puuid:    "test-puuid",
					TeamID:   "456",
					Position: "TOP",
					Role:     "CAPTAIN",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)

			baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetPlayerByPUUID(context.Background(), tt.puuid)

			if tt.wantErr {
				assert.NotNil(t, err)

				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}
				return
			}

			require.Nil(t, err)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetTeamByID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		teamID         string
		httpErr        error
		responseBody   string
		expectedResult Team
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			teamID:       "test-team-id",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			teamID:       "test-team-id",
			statusCode:   http.StatusOK,
			responseBody: `["shouldbeobject"]`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:       "success",
			teamID:     "test-team-id",
			statusCode: http.StatusOK,
			responseBody: `{
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
			}`,
			expectedResult: Team{
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
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetTeamByID(context.Background(), tt.teamID)

			if tt.wantErr {
				assert.NotNil(t, err)

				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}
				return
			}

			require.Nil(t, err)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetTournaments(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		httpErr        error
		responseBody   string
		expectedResult []Tournament
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			statusCode:   http.StatusForbidden,
			responseBody: `{"status":{"status_code":403}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			statusCode:   http.StatusOK,
			responseBody: `{"shouldbearray":true}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			statusCode:     http.StatusOK,
			responseBody:   tournamentsJSON,
			expectedResult: expectedTournaments,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetTournaments(context.Background())

			if tt.wantErr {
				assert.NotNil(t, err)

				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}
				return
			}

			require.Nil(t, err)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetTournamentByTeamID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		teamID         string
		httpErr        error
		responseBody   string
		expectedResult Tournament
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			teamID:       "test-team-id",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			teamID:       "test-team-id",
			statusCode:   http.StatusOK,
			responseBody: `["shouldbeobject"]`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			teamID:         "test-team-id",
			statusCode:     http.StatusOK,
			responseBody:   tournamentJSON,
			expectedResult: expectedTournament,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetTournamentByTeamID(context.Background(), tt.teamID)

			if tt.wantErr {
				assert.NotNil(t, err)

				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}
				return
			}

			require.Nil(t, err)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetTournamentByID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		tournamentID   string
		httpErr        error
		responseBody   string
		expectedResult Tournament
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			tournamentID: "1",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			tournamentID: "1",
			statusCode:   http.StatusOK,
			responseBody: `["shouldbeobject"]`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:           "success",
			tournamentID:   "1",
			statusCode:     http.StatusOK,
			responseBody:   tournamentJSON,
			expectedResult: expectedTournament,
			wantErr:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHTTPClient(mockDoer, slog.Default(), string(regions.PlatformBR1), "apiKey")
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetTournamentByID(context.Background(), tt.tournamentID)

			if tt.wantErr {
				assert.NotNil(t, err)

				if tt.wantRiotErr {
					var rErr *internal.RiotError
					assert.ErrorAs(t, err, &rErr)
					assert.Equal(t, tt.statusCode, rErr.StatusCode)
				}
				return
			}

			require.Nil(t, err)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}
