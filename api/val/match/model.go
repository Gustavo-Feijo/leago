package match

import "github.com/Gustavo-Feijo/leago/internal"

type (
	Match struct {
		MatchInfo    MatchInfo     `json:"matchInfo"`
		Players      []Player      `json:"players"`
		Coaches      []Coach       `json:"coaches"`
		Teams        []Team        `json:"teams"`
		RoundResults []RoundResult `json:"roundResults"`
	}

	MatchInfo struct {
		MatchID            string                  `json:"matchId"`
		MapID              string                  `json:"mapId"`
		GameVersion        string                  `json:"gameVersion"`
		GameLengthMillis   internal.MillisDuration `json:"gameLengthMillis"`
		Region             string                  `json:"region"`
		GameStartMillis    internal.UnixMillisTime `json:"gameStartMillis"`
		ProvisioningFlowID string                  `json:"provisioningFlowId"`
		IsCompleted        bool                    `json:"isCompleted"`
		CustomGameName     string                  `json:"customGameName"`
		QueueID            string                  `json:"queueId"`
		GameMode           string                  `json:"gameMode"`
		IsRanked           bool                    `json:"isRanked"`
		SeasonID           string                  `json:"seasonId"`
		PremierMatchInfo   []PremierMatch          `json:"premierMatchInfo"`
	}

	// Not defined in the docs.
	PremierMatch struct{}

	Player struct {
		PUUID           string      `json:"puuid"`
		GameName        string      `json:"gameName"`
		TagLine         string      `json:"tagLine"`
		TeamID          string      `json:"teamId"`
		PartyID         string      `json:"partyId"`
		CharacterID     string      `json:"characterId"`
		Stats           PlayerStats `json:"stats"`
		CompetitiveTier int         `json:"competitiveTier"`
		IsObserver      bool        `json:"isObserver"`
		PlayerCard      string      `json:"playerCard"`
		PlayerTitle     string      `json:"playerTitle"`
		AccountLevel    int         `json:"accountLevel"`
	}

	PlayerStats struct {
		Score          int                     `json:"score"`
		RoundsPlayed   int                     `json:"roundsPlayed"`
		Kills          int                     `json:"kills"`
		Deaths         int                     `json:"deaths"`
		Assists        int                     `json:"assists"`
		PlaytimeMillis internal.MillisDuration `json:"playtimeMillis"`
		AbilityCasts   AbilityCasts            `json:"abilityCasts"`
	}

	AbilityCasts struct {
		GrenadeCasts  int `json:"grenadeCasts"`
		Ability1Casts int `json:"ability1Casts"`
		Ability2Casts int `json:"ability2Casts"`
		UltimateCasts int `json:"ultimateCasts"`
	}

	Coach struct {
		PUUID  string `json:"puuid"`
		TeamID string `json:"teamId"`
	}

	Team struct {
		TeamID       string `json:"teamId"`
		Won          bool   `json:"won"`
		RoundsPlayed int    `json:"roundsPlayed"`
		RoundsWon    int    `json:"roundsWon"`
		NumPoints    int    `json:"numPoints"`
	}

	RoundResult struct {
		RoundNum              int                `json:"roundNum"`
		RoundResult           string             `json:"roundResult"`
		RoundCeremony         string             `json:"roundCeremony"`
		WinningTeam           string             `json:"winningTeam"`
		WinningTeamRole       string             `json:"winningTeamRole"`
		BombPlanter           string             `json:"bombPlanter"`
		BombDefuser           string             `json:"bombDefuser"`
		PlantRoundTime        int                `json:"plantRoundTime"`
		PlantPlayerLocations  []PlayerLocation   `json:"plantPlayerLocations"`
		PlantLocation         Location           `json:"plantLocation"`
		PlantSite             string             `json:"plantSite"`
		DefuseRoundTime       int                `json:"defuseRoundTime"`
		DefusePlayerLocations []PlayerLocation   `json:"defusePlayerLocations"`
		DefuseLocation        Location           `json:"defuseLocation"`
		PlayerStats           []PlayerRoundStats `json:"playerStats"`
		RoundResultCode       string             `json:"roundResultCode"`
	}

	PlayerLocation struct {
		PUUID       string   `json:"puuid"`
		ViewRadians float64  `json:"viewRadians"`
		Location    Location `json:"location"`
	}

	Location struct {
		X int `json:"x"`
		Y int `json:"y"`
	}

	PlayerRoundStats struct {
		PUUID   string   `json:"puuid"`
		Kills   []Kill   `json:"kills"`
		Damage  []Damage `json:"damage"`
		Score   int      `json:"score"`
		Economy Economy  `json:"economy"`
		Ability Ability  `json:"ability"`
	}

	Kill struct {
		TimeSinceGameStartMillis  internal.MillisDuration `json:"timeSinceGameStartMillis"`
		TimeSinceRoundStartMillis internal.MillisDuration `json:"timeSinceRoundStartMillis"`
		Killer                    string                  `json:"killer"`
		Victim                    string                  `json:"victim"`
		VictimLocation            Location                `json:"victimLocation"`
		Assistants                []string                `json:"assistants"`
		PlayerLocations           []PlayerLocation        `json:"playerLocations"`
		FinishingDamage           FinishingDamage         `json:"finishingDamage"`
	}

	FinishingDamage struct {
		DamageType          string `json:"damageType"`
		DamageItem          string `json:"damageItem"`
		IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
	}

	Damage struct {
		Receiver  string `json:"receiver"`
		Damage    int    `json:"damage"`
		Legshots  int    `json:"legshots"`
		Bodyshots int    `json:"bodyshots"`
		Headshots int    `json:"headshots"`
	}

	Economy struct {
		LoadoutValue int    `json:"loadoutValue"`
		Weapon       string `json:"weapon"`
		Armor        string `json:"armor"`
		Remaining    int    `json:"remaining"`
		Spent        int    `json:"spent"`
	}

	Ability struct {
		GrenadeEffects  string `json:"grenadeEffects"`
		Ability1Effects string `json:"ability1Effects"`
		Ability2Effects string `json:"ability2Effects"`
		UltimateEffects string `json:"ultimateEffects"`
	}
)

type (
	MatchList struct {
		PUUID   string           `json:"puuid"`
		History []MatchListEntry `json:"history"`
	}

	MatchListEntry struct {
		MatchID             string                  `json:"matchId"`
		GameStartTimeMillis internal.UnixMillisTime `json:"gameStartTimeMillis"`
		QueueID             string                  `json:"queueId"`
	}
)

type (
	RecentMatches struct {
		CurrentTime internal.UnixMillisTime `json:"currentTime"`
		MatchIDs    []string                `json:"matchIds"`
	}
)

type Queue string

const (
	QueueCompetitive    Queue = "competitive"
	QueueUnrated        Queue = "unrated"
	QueueSpikerush      Queue = "spikerush"
	QueueTournamentmode Queue = "tournamentmode"
	QueueDeathmatch     Queue = "deathmatch"
	QueueOnefa          Queue = "onefa"
	QueueGgteam         Queue = "ggteam"
	QueueHurm           Queue = "hurm"
	QueueSwiftplay      Queue = "swiftplay"
)
