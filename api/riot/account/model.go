package account

type Account struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

type ActiveRegion struct {
	Puuid       string `json:"puuid"`
	Game        string `json:"game"`
	ActiveShard string `json:"activeShard"`
}

type ActiveShard struct {
	Puuid       string `json:"puuid"`
	Game        string `json:"game"`
	ActiveShard string `json:"activeShard"`
}

type (
	ActiveShardGame  string
	ActiveRegionGame string
)

const (
	ActiveShardValorant ActiveShardGame = "val"
	ActiveShardLOR      ActiveShardGame = "lor"
	ActiveShard2xko     ActiveShardGame = "2xko"

	ActiveRegionLOL ActiveRegionGame = "lol"
	ActiveRegionTFT ActiveRegionGame = "tft"
)
