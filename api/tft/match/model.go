package match

type (
	Match struct {
		Metadata MatchMetadata `json:"metadata"`
		Info     MatchInfo     `json:"info"`
	}

	MatchMetadata struct {
		DataVersion  string   `json:"data_version"`
		MatchID      string   `json:"match_id"`
		Participants []string `json:"participants"`
	}

	MatchInfo struct {
		EndOfGameResult string             `json:"endOfGameResult,omitempty"`
		GameCreation    int64              `json:"gameCreation"`
		GameID          int64              `json:"gameId"`
		GameDatetime    int64              `json:"game_datetime"`
		GameLength      float64            `json:"game_length"`
		GameVersion     string             `json:"game_version"`
		GameVariation   string             `json:"game_variation,omitempty"`
		MapID           int                `json:"mapId"`
		Participants    []MatchParticipant `json:"participants"`
		QueueID         int                `json:"queue_id"`
		QueueIDLegacy   int                `json:"queueId,omitempty"`
		TFTGameType     string             `json:"tft_game_"`
		TFTSetCoreName  string             `json:"tft_set_core_name"`
		TFTSetNumber    int                `json:"tft_set_number"`
	}

	MatchParticipant struct {
		Companion            Companion `json:"companion"`
		GoldLeft             int       `json:"gold_left"`
		LastRound            int       `json:"last_round"`
		Level                int       `json:"level"`
		Placement            int       `json:"placement"`
		PlayersEliminated    int       `json:"players_eliminated"`
		Puuid                string    `json:"puuid"`
		RiotIDGameName       string    `json:"riotIdGameName"`
		RiotIDTagline        string    `json:"riotIdTagline"`
		TimeEliminated       float64   `json:"time_eliminated"`
		TotalDamageToPlayers int       `json:"total_damage_to_players"`
		Traits               []Trait   `json:"traits"`
		Units                []Unit    `json:"units"`
		Win                  bool      `json:"win"`
	}

	Companion struct {
		ContentID string `json:"content_ID"`
		ItemID    int    `json:"item_ID"`
		SkinID    int    `json:"skin_ID"`
		Species   string `json:"species"`
	}

	Trait struct {
		Name        string `json:"name"`
		NumUnits    int    `json:"num_units"`
		Style       int    `json:"style"`
		TierCurrent int    `json:"tier_current"`
		TierTotal   int    `json:"tier_total"`
	}

	Unit struct {
		Items       []int    `json:"items"`
		CharacterID string   `json:"character_id"`
		ItemNames   []string `json:"itemNames"`
		Chosen      string   `json:"chosen,omitempty"`
		Name        string   `json:"name,omitempty"`
		Rarity      int      `json:"rarity"`
		Tier        int      `json:"tier"`
	}
)
