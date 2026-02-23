package account

type (
	Account struct {
		Puuid    string `json:"puuid"`
		GameName string `json:"gameName"`
		TagLine  string `json:"tagLine"`
	}

	ActiveRegion struct {
		Puuid       string `json:"puuid"`
		Game        string `json:"game"`
		ActiveShard string `json:"activeShard"`
	}

	ActiveShard struct {
		Puuid       string `json:"puuid"`
		Game        string `json:"game"`
		ActiveShard string `json:"activeShard"`
	}
)
