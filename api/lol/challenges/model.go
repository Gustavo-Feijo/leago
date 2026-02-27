package challenges

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

type (
	State    string
	Tracking string
	Level    string
	TopLevel string

	ConfigInfo struct {
		ID             int64                        `json:"id"`
		LocalizedNames map[string]map[string]string `json:"localizedNames"`
		State          State                        `json:"state"`
		Tracking       Tracking                     `json:"tracking"`
		StartTimestamp int64                        `json:"startTimestamp"`
		EndTimestamp   int64                        `json:"endTimestamp"`
		Leaderboard    bool                         `json:"leaderboard"`
		Thresholds     map[Level]float64            `json:"thresholds"`
	}

	Leaderboard []struct {
		Puuid    string  `json:"puuid"`
		Value    float64 `json:"value"`
		Position int     `json:"position"`
	}

	PercentileMap    map[int64]LevelPercentiles
	LevelPercentiles map[Level]float64

	PlayerInfo struct {
		Challenges     []PlayerChallenges         `json:"challenges"`
		Preferences    PlayerClientPreferences    `json:"preferences"`
		TotalPoints    ChallengePoints            `json:"totalPoints"`
		CategoryPoints map[string]ChallengePoints `json:"categoryPoints"`
	}

	PlayerChallenges struct {
		Percentiles    float64 `json:"percentile"`
		PlayersInLevel int     `json:"playersInLevel"`
		AchievedTime   int64   `json:"achievedTime"`
		Value          float64 `json:"value"`
		ChallengeID    int64   `json:"challengeId"`
		Level          Level   `json:"level"`
		Position       int     `json:"position"`
	}

	PlayerClientPreferences struct {
		BannerAccent             string   `json:"bannerAccent"`
		Title                    string   `json:"title"`
		ChallengeIds             []string `json:"challengeIds"`
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
