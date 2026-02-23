package champion

type Rotation struct {
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	FreeChampionIds              []int `json:"freeChampionIds"`
}
