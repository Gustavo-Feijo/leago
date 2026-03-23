package tournament

import "github.com/Gustavo-Feijo/leago/internal"

// TournamentCodePayload is the request body for CreateCodes.
type TournamentCodePayload struct {
	SpectatorType SpectatorType `json:"spectatorType"`

	// Valid values from 1-5.
	TeamSize int      `json:"teamSize"`
	PickType PickType `json:"pickType"`

	// AllowedParticipants restricts which PUUIDs may join.
	AllowedParticipants []string `json:"allowedParticipants,omitempty"`

	MapType MapType `json:"mapType"`

	// Optional metadata.
	Metadata      string `json:"metadata"`
	EnoughPlayers bool   `json:"enoughPlayers"`
}

// PutTournamentCodePayload is the request body for UpdateCodes.
type PutTournamentCodePayload struct {
	SpectatorType       SpectatorType `json:"spectatorType"`
	PickType            PickType      `json:"pickType"`
	AllowedParticipants []string      `json:"allowedParticipants,omitempty"`
	MapType             MapType       `json:"mapType"`
}

type (
	// TournamentCode is the full DTO returned by GetCodes.
	TournamentCode struct {
		Code         string           `json:"code"`
		LobbyName    string           `json:"lobbyName"`
		MetaData     string           `json:"metaData"`
		Password     string           `json:"password"` // #nosec G117 - Riot API response field
		TeamSize     int              `json:"teamSize"`
		ProviderID   int              `json:"providerId"`
		PickType     PickType         `json:"pickType"`
		TournamentID int              `json:"tournamentId"`
		ID           int              `json:"id"`
		Region       TournamentRegion `json:"region"`
		Map          MapType          `json:"map"`
		Participants []string         `json:"participants"`
	}
)

type (
	// TournamentGame holds the result of a completed tournament game, returned by GetGamesByCode.
	TournamentGame struct {
		StartTime   int64            `json:"startTime"`
		WinningTeam []TournamentTeam `json:"winningTeam"`
		LosingTeam  []TournamentTeam `json:"losingTeam"`
		ShortCode   string           `json:"shortCode"`
		MetaData    string           `json:"metaData"`
		GameID      int64            `json:"gameId"`

		// If possible move to typed value, the possible values are not listed on the docs.
		GameName string           `json:"gameName"`
		GameType string           `json:"gameType"`
		GameMap  string           `json:"gameMap"`
		GameMode string           `json:"gameMode"`
		Region   TournamentRegion `json:"region"`
	}

	TournamentTeam struct {
		PUUID string `json:"puuid"`
	}
)

type (
	LobbyEventWrapper struct {
		EventList []LobbyEvent `json:"eventList"`
	}

	// LobbyEvent represents a single event in a tournament lobby's history.
	LobbyEvent struct {
		Timestamp internal.LobbyTime `json:"timestamp"`
		// If possible move to typed value, the possible values are not listed on the docs.
		EventType string  `json:"eventType"`
		PUUID     *string `json:"puuid,omitempty"`
	}
)

// ProviderPayload is the request body for CreateProvider.
type ProviderPayload struct {
	Region TournamentRegion `json:"region"`
	URL    string           `json:"url"`
}

// TournamentPayload is the request body for CreateTournament.
type TournamentPayload struct {
	Name     string `json:"name"`
	Provider int    `json:"provider"`
}

type PickType string
type SpectatorType string
type MapType string
type TournamentRegion string

const (
	PickBlind      PickType = "BLIND_PICK"
	PickDraft      PickType = "DRAFT_MODE"
	PickRandom     PickType = "ALL_RANDOM"
	PickTournament PickType = "TOURNAMENT_DRAFT"

	SpectatorNone      SpectatorType = "NONE"
	SpectatorLobbyOnly SpectatorType = "LOBBYONLY"
	SpectatorAll       SpectatorType = "ALL"

	MapSummonersRift MapType = "SUMMONERS_RIFT"
	MapHowlingAbyss  MapType = "HOWLING_ABYSS"

	RegionBR   TournamentRegion = "BR"
	RegionEUNE TournamentRegion = "EUNE"
	RegionEUW  TournamentRegion = "EUW"
	RegionJP   TournamentRegion = "JP"
	RegionLAN  TournamentRegion = "LAN"
	RegionLAS  TournamentRegion = "LAS"
	RegionNA   TournamentRegion = "NA"
	RegionOCE  TournamentRegion = "OCE"
	RegionPBE  TournamentRegion = "PBE"
	RegionRU   TournamentRegion = "RU"
	RegionTR   TournamentRegion = "TR"
	RegionKR   TournamentRegion = "KR"
)
