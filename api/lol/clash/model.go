package clash

import "github.com/Gustavo-Feijo/leago/internal"

// Return a list of registrations of a player to clash tournaments.
type Player struct {
	Puuid    string   `json:"puuid"`
	TeamID   string   `json:"teamId"`
	Position Position `json:"position"`
	Role     Role     `json:"role"`
}

type (
	// Team information for a given clash  tournament.
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
		Puuid    string   `json:"puuid"`
		Position Position `json:"position"`
		Role     Role     `json:"role"`
	}
)

type (
	// Information of a given clash tournament.
	Tournament struct {
		ID               int               `json:"id"`
		ThemeID          int               `json:"themeId"`
		NameKey          string            `json:"nameKey"`
		NameKeySecondary string            `json:"nameKeySecondary"`
		Schedule         []TournamentPhase `json:"schedule"`
	}

	TournamentPhase struct {
		ID               int                     `json:"id"`
		RegistrationTime internal.UnixMillisTime `json:"registrationTime"`
		StartTime        internal.UnixMillisTime `json:"startTime"`
		Cancelled        bool                    `json:"cancelled"`
	}
)

type (
	Position string // The position that the player will play on the team.
	Role     string // Player role, if the team captain or member.
)

const (
	PositionUnselected Position = "UNSELECTED"
	PositionFill       Position = "FILL"
	PositionTop        Position = "TOP"
	PositionJungle     Position = "JUNGLE"
	PositionMiddle     Position = "MIDDLE"
	PositionBottom     Position = "BOTTOM"
	PositionUtility    Position = "UTILITY"

	RoleCaptain Role = "CAPTAIN"
	RoleMember  Role = "MEMBER"
)
