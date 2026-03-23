package account

// Account holds the Riot account identifiers returned by the Account v1 API.
type Account struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

// ActiveRegion reports which Riot region a PUUID is registered in for a specific game.
type ActiveRegion struct {
	Puuid       string `json:"puuid"`
	Game        string `json:"game"`
	ActiveShard string `json:"activeShard"`
}

// ActiveShard reports the active shard (data center cluster) for a PUUID in a specific game.
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
