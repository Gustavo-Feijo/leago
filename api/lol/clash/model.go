package clash

type (
	Player struct {
		Puuid    string `json:"puuid"`
		TeamID   string `json:"teamId"`
		Position string `json:"position"` // UNSELECTED, FILL, TOP, JUNGLE, MIDDLE, BOTTOM, UTILITY
		Role     string `json:"role"`     // CAPTAIN, MEMBER
	}

	PlayersResponse []Player

	Team struct {
		ID           string       `json:"id"`
		TournamentID int          `json:"tournamentId"`
		Name         string       `json:"name"`
		IconID       int          `json:"iconId"`
		Tier         int          `json:"tier"`
		Captain      string       `json:"captain"`
		Abbreviation string       `json:"abbreviation"`
		Players      []TeamPlayer `json:"players"`
	}

	TeamPlayer struct {
		Puuid    string `json:"puuid"`
		Position string `json:"position"` // UNSELECTED, FILL, TOP, JUNGLE, MIDDLE, BOTTOM, UTILITY
		Role     string `json:"role"`     // CAPTAIN, MEMBER
	}

	Tournament struct {
		ID               int               `json:"id"`
		ThemeID          int               `json:"themeId"`
		NameKey          string            `json:"nameKey"`
		NameKeySecondary string            `json:"nameKeySecondary"`
		Schedule         []TournamentPhase `json:"schedule"`
	}

	TournamentPhase struct {
		ID               int   `json:"id"`
		RegistrationTime int64 `json:"registrationTime"`
		StartTime        int64 `json:"startTime"`
		Cancelled        bool  `json:"cancelled"`
	}

	TournamentsResponse []Tournament
)
