package league

type (
	Entry struct {
		Puuid        string      `json:"puuid"`
		LeagueID     string      `json:"leagueId,omitempty"`
		QueueType    string      `json:"queueType"`
		RatedTier    string      `json:"ratedTier,omitempty"`
		RatedRating  int         `json:"ratedRating,omitempty"`
		Tier         string      `json:"tier,omitempty"`
		Rank         string      `json:"rank,omitempty"`
		LeaguePoints int         `json:"leaguePoints,omitempty"`
		Wins         int         `json:"wins"`
		Losses       int         `json:"losses"`
		HotStreak    bool        `json:"hotStreak,omitempty"`
		Veteran      bool        `json:"veteran,omitempty"`
		FreshBlood   bool        `json:"freshBlood,omitempty"`
		Inactive     bool        `json:"inactive,omitempty"`
		MiniSeries   *MiniSeries `json:"miniSeries,omitempty"`
	}

	MiniSeries struct {
		Losses   int    `json:"losses"`
		Progress string `json:"progress"`
		Target   int    `json:"target"`
		Wins     int    `json:"wins"`
	}

	List struct {
		LeagueID string `json:"leagueId"`
		Entries  []Item `json:"entries"`
		Tier     string `json:"tier"`
		Name     string `json:"name"`
		Queue    string `json:"queue"`
	}

	Item struct {
		FreshBlood   bool        `json:"freshBlood"`
		Wins         int         `json:"wins"`
		MiniSeries   *MiniSeries `json:"miniSeries,omitempty"`
		Inactive     bool        `json:"inactive"`
		Veteran      bool        `json:"veteran"`
		HotStreak    bool        `json:"hotStreak"`
		Rank         string      `json:"rank"`
		LeaguePoints int         `json:"leaguePoints"`
		Losses       int         `json:"losses"`
		Puuid        string      `json:"puuid"`
	}

	RatedLadderEntry struct {
		Puuid                        string     `json:"puuid"`
		RatedTier                    LadderTier `json:"ratedTier"`
		RatedRating                  int        `json:"ratedRating"`
		Wins                         int        `json:"wins"`
		PreviousUpdateLadderPosition int        `json:"wreviousUpdateLadderPosition"`
	}

	Queue    string
	Tier     string
	Division string

	LadderQueue string
	LadderTier  string
)

const (
	QueueRankedTFT         Queue = "RANKED_TFT"
	QueueRankedTFTDoubleUP Queue = "RANKED_TFT_DOUBLE_UP"

	TierDiamond  Tier = "DIAMOND"
	TierEmerald  Tier = "EMERALD"
	TierPlatinum Tier = "PLATINUM"
	TierGold     Tier = "GOLD"
	TierSilver   Tier = "SILVER"
	TierBronze   Tier = "BRONZE"
	TierIron     Tier = "IRON"

	DivisionI   Division = "I"
	DivisionII  Division = "II"
	DivisionIII Division = "III"
	DivisionIV  Division = "IV"

	LadderQueueRankedTFTDoubleUP LadderQueue = "RANKED_TFT_DOUBLE_UP"

	LadderOrange LadderTier = "ORANGE"
	LadderPurple LadderTier = "PURPLE"
	LadderBlue   LadderTier = "BLUE"
	LadderGreen  LadderTier = "GREEN"
	LadderGray   LadderTier = "GRAY"
)
