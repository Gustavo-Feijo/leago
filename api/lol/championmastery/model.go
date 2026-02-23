package championmastery

type (
	MasteryList []Mastery

	Mastery struct {
		Puuid                        string               `json:"puuid"`
		ChampionPointsUntilNextLevel int64                `json:"championPointsUntilNextLevel"`
		ChestGranted                 bool                 `json:"chestGranted"`
		ChampionID                   int64                `json:"championId"`
		LastPlayTime                 int64                `json:"lastPlayTime"`
		ChampionLevel                int                  `json:"championLevel"`
		ChampionPoints               int                  `json:"championPoints"`
		ChampionPointsSinceLastLevel int64                `json:"championPointsSinceLastLevel"`
		MarkRequiredForNextLevel     int                  `json:"markRequiredForNextLevel"`
		ChampionSeasonMilestone      int                  `json:"championSeasonMilestone"`
		NextSeasonMilestone          NextSeasonMilestones `json:"nextSeasonMilestone"`
		TokensEarned                 int                  `json:"tokensEarned"`
		MilestoneGrades              []string             `json:"milestoneGrades"`
	}

	NextSeasonMilestones struct {
		RequireGradeCounts map[string]int `json:"requireGradeCounts"`
		RewardMarks        int            `json:"rewardMarks"`
		Bonus              bool           `json:"bonus"`
		RewardConfig       RewardConfig   `json:"rewardConfig"`
	}

	RewardConfig struct {
		RewardValue   string `json:"rewardValue"`
		RewardType    string `json:"rewardType"`
		MaximumReward int    `json:"maximumReward"`
	}

	MasteryScore int
)
