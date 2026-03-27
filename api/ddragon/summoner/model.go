package summoner

type (
	SummonerResponse struct {
		Type    string              `json:"type"`
		Version string              `json:"version"`
		Data    map[string]Summoner `json:"data"`
	}

	Summoner struct {
		ID            string         `json:"id"`
		Name          string         `json:"name"`
		Description   string         `json:"description"`
		Tooltip       string         `json:"tooltip"`
		MaxRank       int            `json:"maxrank"`
		Cooldown      []float64      `json:"cooldown"`
		CooldownBurn  string         `json:"cooldownBurn"`
		Cost          []int          `json:"cost"`
		CostBurn      string         `json:"costBurn"`
		DataValues    map[string]any `json:"datavalues"`
		Effect        [][]float64    `json:"effect"`
		EffectBurn    []*string      `json:"effectBurn"`
		Vars          []any          `json:"vars"`
		Key           string         `json:"key"`
		SummonerLevel int            `json:"summonerLevel"`
		Modes         []string       `json:"modes"`
		CostType      string         `json:"costType"`
		MaxAmmo       string         `json:"maxammo"`
		Range         []int          `json:"range"`
		RangeBurn     string         `json:"rangeBurn"`
		Image         Image          `json:"image"`
		Resource      string         `json:"resource"`
	}

	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	}
)
