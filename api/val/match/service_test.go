package match

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
	expectedMatch = Match{
		MatchInfo: MatchInfo{
			MatchID:            "test-match",
			MapID:              "ascent",
			GameVersion:        "1.0",
			GameLengthMillis:   internal.MillisDuration{Duration: 900000 * time.Millisecond},
			Region:             "br",
			GameStartMillis:    internal.UnixMillisTime{Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
			ProvisioningFlowID: "flow",
			IsCompleted:        true,
			QueueID:            "competitive",
			GameMode:           "bomb",
			IsRanked:           true,
			SeasonID:           "season-1",
			PremierMatchInfo:   []PremierMatch{},
		},

		Players: []Player{
			{
				PUUID:       "player1",
				GameName:    "player",
				TagLine:     "BR1",
				TeamID:      "Blue",
				PartyID:     "party1",
				CharacterID: "agent1",
				Stats: PlayerStats{
					Score:          100,
					RoundsPlayed:   1,
					Kills:          1,
					Deaths:         0,
					Assists:        0,
					PlaytimeMillis: internal.MillisDuration{Duration: 900000 * time.Millisecond},
					AbilityCasts: AbilityCasts{
						GrenadeCasts:  0,
						Ability1Casts: 0,
						Ability2Casts: 0,
						UltimateCasts: 0,
					},
				},
				CompetitiveTier: 10,
				IsObserver:      false,
				PlayerCard:      "card",
				PlayerTitle:     "title",
				AccountLevel:    1,
			},
		},

		Coaches: []Coach{},

		Teams: []Team{
			{
				TeamID:       "Blue",
				Won:          true,
				RoundsPlayed: 1,
				RoundsWon:    1,
				NumPoints:    1,
			},
		},

		RoundResults: []RoundResult{
			{
				RoundNum:              1,
				RoundResult:           "Eliminated",
				RoundCeremony:         "CeremonyDefault",
				WinningTeam:           "Blue",
				WinningTeamRole:       "Attack",
				BombPlanter:           "",
				BombDefuser:           "",
				PlantRoundTime:        0,
				PlantLocation:         Location{X: 0, Y: 0},
				PlantSite:             "A",
				DefuseRoundTime:       0,
				DefuseLocation:        Location{X: 0, Y: 0},
				PlantPlayerLocations:  []PlayerLocation{},
				DefusePlayerLocations: []PlayerLocation{},
				PlayerStats: []PlayerRoundStats{
					{
						PUUID:  "player1",
						Kills:  []Kill{},
						Damage: []Damage{},
						Score:  100,
						Economy: Economy{
							LoadoutValue: 1000,
							Weapon:       "rifle",
							Armor:        "light",
							Remaining:    0,
							Spent:        1000,
						},
						Ability: Ability{
							GrenadeEffects:  "",
							Ability1Effects: "",
							Ability2Effects: "",
							UltimateEffects: "",
						},
					},
				},

				RoundResultCode: "Elimination",
			},
		},
	}

	matchJSON = `{
  "matchInfo": {
    "matchId": "test-match",
    "mapId": "ascent",
    "gameVersion": "1.0",
    "gameLengthMillis": 900000,
    "region": "br",
    "gameStartMillis": 1767225600000,
    "provisioningFlowId": "flow",
    "isCompleted": true,
    "customGameName": "",
    "queueId": "competitive",
    "gameMode": "bomb",
    "isRanked": true,
    "seasonId": "season-1",
    "premierMatchInfo": []
  },
  "players": [
    {
      "puuid": "player1",
      "gameName": "player",
      "tagLine": "BR1",
      "teamId": "Blue",
      "partyId": "party1",
      "characterId": "agent1",
      "stats": {
        "score": 100,
        "roundsPlayed": 1,
        "kills": 1,
        "deaths": 0,
        "assists": 0,
        "playtimeMillis": 900000,
        "abilityCasts": {
          "grenadeCasts": 0,
          "ability1Casts": 0,
          "ability2Casts": 0,
          "ultimateCasts": 0
        }
      },
      "competitiveTier": 10,
      "isObserver": false,
      "playerCard": "card",
      "playerTitle": "title",
      "accountLevel": 1
    }
  ],
  "coaches": [],
  "teams": [
    {
      "teamId": "Blue",
      "won": true,
      "roundsPlayed": 1,
      "roundsWon": 1,
      "numPoints": 1
    }
  ],
  "roundResults": [
    {
      "roundNum": 1,
      "roundResult": "Eliminated",
      "roundCeremony": "CeremonyDefault",
      "winningTeam": "Blue",
      "winningTeamRole": "Attack",
      "bombPlanter": "",
      "bombDefuser": "",
      "plantRoundTime": 0,
      "plantPlayerLocations": [],
      "plantLocation": { "x": 0, "y": 0 },
      "plantSite": "A",
      "defuseRoundTime": 0,
      "defusePlayerLocations": [],
      "defuseLocation": { "x": 0, "y": 0 },
      "playerStats": [
        {
          "puuid": "player1",
          "kills": [],
          "damage": [],
          "score": 100,
          "economy": {
            "loadoutValue": 1000,
            "weapon": "rifle",
            "armor": "light",
            "remaining": 0,
            "spent": 1000
          },
          "ability": {
            "grenadeEffects": "",
            "ability1Effects": "",
            "ability2Effects": "",
            "ultimateEffects": ""
          }
        }
      ],
      "roundResultCode": "Elimination"
    }
  ]
}`

	expectedMatchlist = MatchList{
		PUUID: "player1",
		History: []MatchListEntry{
			{
				MatchID:             "match-123",
				GameStartTimeMillis: internal.UnixMillisTime{Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
				QueueID:             "competitive",
			},
			{
				MatchID:             "match-456",
				GameStartTimeMillis: internal.UnixMillisTime{Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
				QueueID:             "unrated",
			},
		},
	}

	matchlistJSON = `
	{
  "puuid": "player1",
  "history": [
    {
      "matchId": "match-123",
      "gameStartTimeMillis": 1767225600000,
      "queueId": "competitive"
    },
    {
      "matchId": "match-456",
      "gameStartTimeMillis": 1767225600000,
      "queueId": "unrated"
    }
  ]
}`

	expectedRecentMatches = RecentMatches{
		CurrentTime: internal.UnixMillisTime{Time: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)},
		MatchIDs:    []string{"match1", "match2"},
	}

	recentmatchesJSON = `
	{
	   "currentTime":1767225600000,
	   "matchIds":[
	      "match1",
	      "match2"
	   ]
	}
	`
)

func newTestRegionClient(statusCode int, responseBody string) (*RegionClient, *mock.Doer) {
	mockDoer := mock.NewDefaultDoer(statusCode, responseBody)
	baseClient := internal.NewHTTPClient(string(regions.RegionAmericas), "apiKey", internal.WithHTTP(mockDoer))

	return NewRegionClient(baseClient), mockDoer
}

func TestGetMatchByID(t *testing.T) {
	tests := []struct {
		name string

		matchID string

		statusCode   int
		responseBody string

		expectedPath string

		expectedResult Match

		wantErr     bool
		wantRiotErr bool
	}{
		{
			name: "riot error",

			matchID: "nonexistent",

			statusCode:   http.StatusNotFound,
			responseBody: `{"status":{"status_code":404}}`,

			wantErr:     true,
			wantRiotErr: true,
		},
		{
			name: "invalid json",

			matchID: "badjsonmatch",

			statusCode:   http.StatusOK,
			responseBody: `{"invalid json,,,,::"}`,

			wantErr: true,
		},
		{
			name: "success",

			matchID: "BR1_123",

			statusCode:   http.StatusOK,
			responseBody: matchJSON,

			expectedPath: "/val/match/v1/matches/BR1_123",

			expectedResult: expectedMatch,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rc, mockDoer := newTestRegionClient(tt.statusCode, tt.responseBody)

			resp, err := rc.GetMatchByID(context.Background(), "BR1_123")

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

func TestGetMatchesByPUUID(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, matchlistJSON)

	resp, err := rc.GetMatchesByPUUID(
		context.Background(),
		"testpuuid",
	)

	require.NoError(t, err)

	assert.Equal(t, "/val/match/v1/matchlists/by-puuid/testpuuid", mockDoer.CapturedReq.URL.Path)

	assert.Equal(t, expectedMatchlist, resp)
}

func TestGetRecentMatchesByQueue(t *testing.T) {
	rc, mockDoer := newTestRegionClient(http.StatusOK, recentmatchesJSON)

	resp, err := rc.GetRecentMatchesByQueue(
		context.Background(),
		QueueCompetitive,
	)

	require.NoError(t, err)

	assert.Equal(t, "/val/match/v1/recent-matches/by-queue/competitive", mockDoer.CapturedReq.URL.Path)

	assert.Equal(t, expectedRecentMatches, resp)
}
