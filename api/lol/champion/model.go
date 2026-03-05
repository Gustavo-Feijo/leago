package champion

// Current free champion rotation.
type Rotation struct {
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
	FreeChampionIDsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	FreeChampionIDs              []int `json:"freeChampionIds"`
}
