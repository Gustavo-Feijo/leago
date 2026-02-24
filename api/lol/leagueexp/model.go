package leagueexp

type (
	LeagueResponse []LeagueEntry

	LeagueEntry struct {
		LeagueID     string      `json:"leagueId"`
		SummonerID   string      `json:"summonerId"`
		PUUID        string      `json:"puuid"`
		QueueType    string      `json:"queueType"`
		Tier         string      `json:"tier"`
		Rank         string      `json:"rank"`
		LeaguePoints int         `json:"leaguePoints"`
		Wins         int         `json:"wins"`
		Losses       int         `json:"losses"`
		HotStreak    bool        `json:"hotStreak"`
		Veteran      bool        `json:"veteran"`
		FreshBlood   bool        `json:"freshBlood"`
		Inactive     bool        `json:"inactive"`
		MiniSeries   *MiniSeries `json:"miniSeries,omitempty"`
	}

	MiniSeries struct {
		Losses   int    `json:"losses"`
		Progress string `json:"progress"`
		Target   int    `json:"target"`
		Wins     int    `json:"wins"`
	}

	Queue    string
	Tier     string
	Division string
)

const (
	QueueRankedSolo   Queue = "RANKED_SOLO_5x5"
	QueueRankedTFT    Queue = "RANKED_TFT"
	QueueRankedFlexSR Queue = "RANKED_FLEX_SR"
	QueueRankedFlexTT Queue = "RANKED_FLEX_TT"

	TierChallenger  Tier = "CHALLENGER"
	TierGrandmaster Tier = "GRANDMASTER"
	TierMaster      Tier = "MASTER"
	TierDiamond     Tier = "DIAMOND"
	TierEmerald     Tier = "EMERALD"
	TierPlatinum    Tier = "PLATINUM"
	TierGold        Tier = "GOLD"
	TierSilver      Tier = "SILVER"
	TierBronze      Tier = "BRONZE"
	TierIron        Tier = "IRON"

	DivisionI   Division = "I"
	DivisionII  Division = "II"
	DivisionIII Division = "III"
	DivisionIV  Division = "IV"
)
