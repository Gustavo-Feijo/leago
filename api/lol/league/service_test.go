package league

import (
	"context"
	"leago/internal"
	"leago/internal/mock"
	"leago/regions"
	"log/slog"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetChallengerLeague(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		queue          Queue
		httpErr        error
		responseBody   string
		expectedResult RawLeague
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			queue:        QueueRankedSolo,
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			queue:        QueueRankedSolo,
			statusCode:   http.StatusOK,
			responseBody: `["invalid"]`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:       "success",
			queue:      QueueRankedSolo,
			statusCode: http.StatusOK,
			responseBody: `{
				"leagueId":"league-1",
				"entries":[{
					"puuid":"puuid-1",
					"rank":"I",
					"leaguePoints":1000,
					"wins":200,
					"losses":100,
					"hotStreak":true,
					"veteran":true,
					"freshBlood":false,
					"inactive":false
				}],
				"tier":"CHALLENGER",
				"name":"test league",
				"queue":"RANKED_SOLO_5x5"
			}`,
			expectedResult: RawLeague{
				LeagueID: "league-1",
				Entries: []RawEloEntry{
					{
						PUUID:        "puuid-1",
						Rank:         "I",
						LeaguePoints: 1000,
						Wins:         200,
						Losses:       100,
						HotStreak:    true,
						Veteran:      true,
						FreshBlood:   false,
						Inactive:     false,
						MiniSeries:   nil,
					},
				},
				Tier:  "CHALLENGER",
				Name:  "test league",
				Queue: "RANKED_SOLO_5x5",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHttpClient(
				mockDoer,
				slog.Default(),
				string(regions.PlatformBR1),
				"apiKey",
			)
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetChallengerLeague(context.Background(), tt.queue)

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

func TestGetGrandmasterLeague(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		queue          Queue
		httpErr        error
		responseBody   string
		expectedResult RawLeague
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			queue:        QueueRankedSolo,
			statusCode:   http.StatusUnauthorized,
			responseBody: `{"status":{"status_code":401}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:       "success",
			queue:      QueueRankedFlexSR,
			statusCode: http.StatusOK,
			responseBody: `{
				"leagueId":"league-2",
				"entries":[{
					"puuid":"puuid-2",
					"rank":"I",
					"leaguePoints":500,
					"wins":150,
					"losses":80,
					"hotStreak":false,
					"veteran":false,
					"freshBlood":true,
					"inactive":false
				}],
				"tier":"GRANDMASTER",
				"name":"grandmaster league",
				"queue":"RANKED_FLEX_SR"
			}`,
			expectedResult: RawLeague{
				LeagueID: "league-2",
				Entries: []RawEloEntry{
					{
						PUUID:        "puuid-2",
						Rank:         "I",
						LeaguePoints: 500,
						Wins:         150,
						Losses:       80,
						HotStreak:    false,
						Veteran:      false,
						FreshBlood:   true,
						Inactive:     false,
						MiniSeries:   nil,
					},
				},
				Tier:  "GRANDMASTER",
				Name:  "grandmaster league",
				Queue: "RANKED_FLEX_SR",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHttpClient(
				mockDoer,
				slog.Default(),
				string(regions.PlatformBR1),
				"apiKey",
			)
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetGrandmasterLeague(context.Background(), tt.queue)

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

func TestGetMasterLeague(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		queue          Queue
		httpErr        error
		responseBody   string
		expectedResult RawLeague
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			queue:        QueueRankedSolo,
			statusCode:   http.StatusForbidden,
			responseBody: `{"status":{"status_code":403}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:       "success",
			queue:      QueueRankedSolo,
			statusCode: http.StatusOK,
			responseBody: `{
				"leagueId":"league-3",
				"entries":[{
					"puuid":"puuid-3",
					"rank":"I",
					"leaguePoints":100,
					"wins":50,
					"losses":40,
					"hotStreak":false,
					"veteran":true,
					"freshBlood":false,
					"inactive":false
				}],
				"tier":"MASTER",
				"name":"master league",
				"queue":"RANKED_SOLO_5x5"
			}`,
			expectedResult: RawLeague{
				LeagueID: "league-3",
				Entries: []RawEloEntry{
					{
						PUUID:        "puuid-3",
						Rank:         "I",
						LeaguePoints: 100,
						Wins:         50,
						Losses:       40,
						HotStreak:    false,
						Veteran:      true,
						FreshBlood:   false,
						Inactive:     false,
						MiniSeries:   nil,
					}},
				Tier:  "MASTER",
				Name:  "master league",
				Queue: "RANKED_SOLO_5x5",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHttpClient(
				mockDoer,
				slog.Default(),
				string(regions.PlatformBR1),
				"apiKey",
			)
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetMasterLeague(context.Background(), tt.queue)

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

func TestGetLeagueEntries(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		queue          Queue
		tier           Tier
		division       Division
		httpErr        error
		responseBody   string
		expectedResult []Entry
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			queue:        QueueRankedSolo,
			tier:         TierGold,
			division:     DivisionI,
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			queue:        QueueRankedSolo,
			tier:         TierGold,
			division:     DivisionI,
			statusCode:   http.StatusOK,
			responseBody: `{"shouldbearray":true}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:       "success without mini series",
			queue:      QueueRankedSolo,
			tier:       TierGold,
			division:   DivisionI,
			statusCode: http.StatusOK,
			responseBody: `[{
				"leagueId":"league-1",
				"summonerId":"summoner-1",
				"puuid":"puuid-1",
				"queueType":"RANKED_SOLO_5x5",
				"tier":"GOLD",
				"rank":"I",
				"leaguePoints":75,
				"wins":10,
				"losses":5,
				"hotStreak":true,
				"veteran":false,
				"freshBlood":true,
				"inactive":false
			}]`,
			expectedResult: []Entry{
				{
					LeagueID:     "league-1",
					SummonerID:   "summoner-1",
					PUUID:        "puuid-1",
					QueueType:    "RANKED_SOLO_5x5",
					Tier:         "GOLD",
					Rank:         "I",
					LeaguePoints: 75,
					Wins:         10,
					Losses:       5,
					HotStreak:    true,
					Veteran:      false,
					FreshBlood:   true,
					Inactive:     false,
					MiniSeries:   nil,
				},
			},
			wantErr: false,
		},
		{
			name:       "success with mini series",
			queue:      QueueRankedSolo,
			tier:       TierGold,
			division:   DivisionI,
			statusCode: http.StatusOK,
			responseBody: `[{
				"leagueId":"league-2",
				"summonerId":"summoner-2",
				"puuid":"puuid-2",
				"queueType":"RANKED_SOLO_5x5",
				"tier":"GOLD",
				"rank":"I",
				"leaguePoints":100,
				"wins":20,
				"losses":10,
				"hotStreak":false,
				"veteran":true,
				"freshBlood":false,
				"inactive":false,
				"miniSeries":{
					"losses":1,
					"progress":"WLN",
					"target":3,
					"wins":1
				}
			}]`,
			expectedResult: []Entry{
				{
					LeagueID:     "league-2",
					SummonerID:   "summoner-2",
					PUUID:        "puuid-2",
					QueueType:    "RANKED_SOLO_5x5",
					Tier:         "GOLD",
					Rank:         "I",
					LeaguePoints: 100,
					Wins:         20,
					Losses:       10,
					HotStreak:    false,
					Veteran:      true,
					FreshBlood:   false,
					Inactive:     false,
					MiniSeries: &MiniSeries{
						Losses:   1,
						Progress: "WLN",
						Target:   3,
						Wins:     1,
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHttpClient(
				mockDoer,
				slog.Default(),
				string(regions.PlatformBR1),
				"apiKey",
			)
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetLeagueEntries(
				context.Background(),
				tt.queue,
				tt.tier,
				tt.division,
				[]GetLeagueOption{WithPage(1)},
			)

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
			require.NotNil(t, resp)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetLeagueEntriesByPUUID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		puuid          string
		httpErr        error
		responseBody   string
		expectedResult []Entry
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			puuid:        "puuid-1",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			puuid:        "puuid-1",
			statusCode:   http.StatusOK,
			responseBody: `{"shouldbearray":true}`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:       "success",
			puuid:      "puuid-1",
			statusCode: http.StatusOK,
			responseBody: `[{
				"leagueId":"league-1",
				"summonerId":"summoner-1",
				"puuid":"puuid-1",
				"queueType":"RANKED_SOLO_5x5",
				"tier":"GOLD",
				"rank":"II",
				"leaguePoints":50,
				"wins":30,
				"losses":20,
				"hotStreak":false,
				"veteran":false,
				"freshBlood":false,
				"inactive":false
			}]`,
			expectedResult: []Entry{
				{
					LeagueID:     "league-1",
					SummonerID:   "summoner-1",
					PUUID:        "puuid-1",
					QueueType:    "RANKED_SOLO_5x5",
					Tier:         "GOLD",
					Rank:         "II",
					LeaguePoints: 50,
					Wins:         30,
					Losses:       20,
					HotStreak:    false,
					Veteran:      false,
					FreshBlood:   false,
					Inactive:     false,
					MiniSeries:   nil,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHttpClient(
				mockDoer,
				slog.Default(),
				string(regions.PlatformBR1),
				"apiKey",
			)
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetLeagueEntriesByPUUID(context.Background(), tt.puuid)

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
			require.NotNil(t, resp)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}

func TestGetLeagueByID(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		leagueID       string
		httpErr        error
		responseBody   string
		expectedResult RawLeague
		wantErr        bool
		wantRiotErr    bool
	}{
		{
			name:         "riot error",
			leagueID:     "league-1",
			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,
			wantErr:      true,
			wantRiotErr:  true,
		},
		{
			name:         "unmatched json",
			leagueID:     "league-1",
			statusCode:   http.StatusOK,
			responseBody: `[{"shouldntbearrayofobj":true}]`,
			wantErr:      true,
			wantRiotErr:  false,
		},
		{
			name:       "success without mini series",
			leagueID:   "league-1",
			statusCode: http.StatusOK,
			responseBody: `{
        		"leagueId":"league-1",
        		"tier":"PLATINUM",
        		"name":"Test League",
        		"queue":"RANKED_SOLO_5x5",
        		"entries":[{
        		    "puuid":"puuid-1",
        		    "rank":"III",
        		    "leaguePoints":25,
        		    "wins":60,
        		    "losses":55,
        		    "hotStreak":false,
        		    "veteran":true,
        		    "freshBlood":false,
        		    "inactive":false
        		}]
   			}`,
			expectedResult: RawLeague{
				LeagueID: "league-1",
				Tier:     "PLATINUM",
				Name:     "Test League",
				Queue:    "RANKED_SOLO_5x5",
				Entries: []RawEloEntry{
					{
						PUUID:        "puuid-1",
						Rank:         "III",
						LeaguePoints: 25,
						Wins:         60,
						Losses:       55,
						HotStreak:    false,
						Veteran:      true,
						FreshBlood:   false,
						Inactive:     false,
						MiniSeries:   nil,
					},
				},
			},
			wantErr: false,
		},
		{
			name:       "success with mini series",
			leagueID:   "league-2",
			statusCode: http.StatusOK,
			responseBody: `{
        		"leagueId":"league-2",
        		"tier":"SILVER",
        		"name":"Silver League",
        		"queue":"RANKED_SOLO_5x5",
        		"entries":[{
        		    "puuid":"puuid-2",
        		    "rank":"I",
        		    "leaguePoints":100,
        		    "wins":40,
        		    "losses":35,
        		    "hotStreak":true,
        		    "veteran":false,
        		    "freshBlood":false,
        		    "inactive":false,
        		    "miniSeries":{
        		        "losses":0,
        		        "progress":"WNN",
        		        "target":3,
        		        "wins":1
        		    }
        		}]
   			}`,
			expectedResult: RawLeague{
				LeagueID: "league-2",
				Tier:     "SILVER",
				Name:     "Silver League",
				Queue:    "RANKED_SOLO_5x5",
				Entries: []RawEloEntry{
					{
						PUUID:        "puuid-2",
						Rank:         "I",
						LeaguePoints: 100,
						Wins:         40,
						Losses:       35,
						HotStreak:    true,
						Veteran:      false,
						FreshBlood:   false,
						Inactive:     false,
						MiniSeries: &MiniSeries{
							Losses:   0,
							Progress: "WNN",
							Target:   3,
							Wins:     1,
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDoer := mock.NewDefaultDoer(tt.statusCode, tt.responseBody, tt.httpErr)
			baseClient := internal.NewHttpClient(
				mockDoer,
				slog.Default(),
				string(regions.PlatformBR1),
				"apiKey",
			)
			pc := NewPlatformClient(baseClient)
			resp, err := pc.GetLeagueByID(context.Background(), tt.leagueID)

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
			require.NotNil(t, resp)
			assert.Equal(t, tt.expectedResult, resp)
		})
	}
}
