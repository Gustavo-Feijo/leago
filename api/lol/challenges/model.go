package challenges

import "github.com/Gustavo-Feijo/leago/internal"

// List of Basic challenge configuration info (Includes translations).
type ConfigInfo struct {
	ID             int64                        `json:"id"`
	LocalizedNames map[string]map[string]string `json:"localizedNames"`
	State          State                        `json:"state"`
	Tracking       Tracking                     `json:"tracking"`
	StartTimestamp internal.UnixMillisTime      `json:"startTimestamp"`
	EndTimestamp   internal.UnixMillisTime      `json:"endTimestamp"`
	Leaderboard    bool                         `json:"leaderboard"`
	Thresholds     map[Level]float64            `json:"thresholds"`
}

type (
	// Map representing the percentile of players to achieve a given level in a challenge.
	PercentileMap    map[int64]LevelPercentiles
	LevelPercentiles map[Level]float64
)

// Leaderboard of top players on a given challenge.
type Leaderboard []struct {
	Puuid    string  `json:"puuid"`
	Value    float64 `json:"value"`
	Position int     `json:"position"`
}

type (
	// Player information with their challenges progression.
	PlayerInfo struct {
		Challenges     []PlayerChallenges         `json:"challenges"`
		Preferences    PlayerClientPreferences    `json:"preferences"`
		TotalPoints    ChallengePoints            `json:"totalPoints"`
		CategoryPoints map[string]ChallengePoints `json:"categoryPoints"`
	}

	PlayerChallenges struct {
		Percentiles    float64                 `json:"percentile"`
		PlayersInLevel int                     `json:"playersInLevel"`
		AchievedTime   internal.UnixMillisTime `json:"achievedTime"`
		Value          float64                 `json:"value"`
		ChallengeID    int64                   `json:"challengeId"`
		Level          Level                   `json:"level"`
		Position       int                     `json:"position"`
	}

	PlayerClientPreferences struct {
		BannerAccent             string   `json:"bannerAccent"`
		Title                    string   `json:"title"`
		ChallengeIDs             []string `json:"challengeIds"`
		CrestBorder              string   `json:"crestBorder"`
		PrestigeCrestBorderLevel int      `json:"prestigeCrestBorderLevel"`
	}

	ChallengePoints struct {
		Level      string  `json:"level"`
		Current    int64   `json:"current"`
		Max        int64   `json:"max"`
		Percentile float64 `json:"percentile"`
	}
)

type (
	State    string // Current state of a given challenge.
	Tracking string // How the challenge is tracked, if it is seasonal or lifetime.

	Level    string // Challenge level.
	TopLevel string // Challenge level, only including the top rated levels.
)

const (
	StateDisabled State = "DISABLED"
	StateHidden   State = "HIDDEN"
	StateEnabled  State = "ENABLED"
	StateArchived State = "ARCHIVED"

	TrackingLifetime Tracking = "LIFETIME"
	TrackingSeason   Tracking = "SEASON"

	LevelNone                      Level = "NONE"
	LevelIron                      Level = "IRON"
	LevelBronze                    Level = "BRONZE"
	LevelSilver                    Level = "SILVER"
	LevelGold                      Level = "GOLD"
	LevelPlatinum                  Level = "PLATINUM"
	LevelDiamond                   Level = "DIAMOND"
	LevelMaster                    Level = "MASTER"
	LevelGrandmaster               Level = "GRANDMASTER"
	LevelChallenger                Level = "CHALLENGER"
	LevelHighestNotLeaderboardOnly Level = "HIGHEST_NOT_LEADERBOARD_ONLY"
	LevelHighest                   Level = "HIGHEST"
	LevelLowest                    Level = "LOWEST"

	// Leaderboards endpoint only accept those.
	TopLevelMaster      TopLevel = "MASTER"
	TopLevelGrandmaster TopLevel = "GRANDMASTER"
	TopLevelChallenger  TopLevel = "CHALLENGER"
)
