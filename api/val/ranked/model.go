package ranked

type Leaderboard struct {
	Shard        string   `json:"shard"`
	ActID        string   `json:"actId"`
	TotalPlayers int64    `json:"totalPlayers"`
	Players      []Player `json:"players"`
}

type Player struct {
	PUUID           *string `json:"puuid,omitempty"`
	GameName        *string `json:"gameName,omitempty"`
	TagLine         *string `json:"tagLine,omitempty"`
	LeaderboardRank int64   `json:"leaderboardRank"`
	RankedRating    int64   `json:"rankedRating"`
	NumberOfWins    int64   `json:"numberOfWins"`
}
