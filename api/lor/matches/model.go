package matches

import "time"

type (
	Match struct {
		Metadata Metadata `json:"metadata"`
		Info     Info     `json:"info"`
	}

	Metadata struct {
		DataVersion  string   `json:"data_version"`
		MatchID      string   `json:"match_id"`
		Participants []string `json:"participants"`
	}

	Info struct {
		GameMode         GameMode   `json:"game_mode"`
		GameType         GameType   `json:"game_type"`
		GameStartTimeUTC time.Time  `json:"game_start_time_utc"`
		GameVersion      string     `json:"game_version"`
		GameFormat       GameFormat `json:"game_format"`
		Players          []Player   `json:"players"`
		TotalTurnCount   int        `json:"total_turn_count"`
	}

	Player struct {
		PUUID       string   `json:"puuid"`
		DeckID      string   `json:"deck_id"`
		DeckCode    string   `json:"deck_code"`
		Factions    []string `json:"factions"`
		GameOutcome string   `json:"game_outcome"`
		OrderOfPlay int      `json:"order_of_play"`
	}

	GameMode   string
	GameType   string
	GameFormat string
)

const (
	GameModeConstructed        GameMode = "Constructed"
	GameModeExpeditions        GameMode = "Expeditions"
	GameModeTutorial           GameMode = "Tutorial"
	GameModeThePathOfChampions GameMode = "ThePathOfChampions"

	GameTypeRanked           GameType = "Ranked"
	GameTypeNormal           GameType = "Normal"
	GameTypeAI               GameType = "AI"
	GameTypeTutorial         GameType = "Tutorial"
	GameTypeVanillaTrial     GameType = "VanillaTrial"
	GameTypeSingleton        GameType = "Singleton"
	GameTypeStandardGauntlet GameType = "StandardGauntlet"

	GameFormatStandard GameFormat = "standard"
	GameFormatEternal  GameFormat = "eternal"
)
