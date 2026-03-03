package ranked

type (
	Leaderboard struct {
		Players []Player `json:"players"`
	}

	Player struct {
		Name string  `json:"name"`
		Rank int     `json:"rank"`
		LP   float64 `json:"lp"` // Weirdly docs and request inside the developer portal return as int64, but the API truly return float.
	}
)
