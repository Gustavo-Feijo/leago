package match

type (
	Timeline struct {
		Metadata MetadataTimeLine `json:"metadata"`
		Info     InfoTimeLine     `json:"info"`
	}

	MetadataTimeLine struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	}

	InfoTimeLine struct {
		EndOfGameResult string                `json:"endOfGameResult"`
		FrameInterval   int64                 `json:"frameInterval"`
		GameID          int64                 `json:"gameId"`
		Participants    []ParticipantTimeLine `json:"participants"`
		Frames          []FrameTimeLine       `json:"frames"`
	}

	ParticipantTimeLine struct {
		ParticipantID int    `json:"participantId"`
		PUUID         string `json:"puuid"`
	}

	FrameTimeLine struct {
		Events            []EventTimeLine   `json:"events"`
		ParticipantFrames ParticipantFrames `json:"participantFrames"`
		Timestamp         int               `json:"timestamp"`
	}

	EventTimeLine struct {
		Timestamp     int64  `json:"timestamp"`
		RealTimestamp int64  `json:"realTimestamp"`
		Type          string `json:"type"`
	}

	ParticipantFrames map[string]ParticipantFrame

	ParticipantFrame struct {
		ChampionStats            ChampionStats `json:"championStats"`
		CurrentGold              int           `json:"currentGold"`
		DamageStats              DamageStats   `json:"damageStats"`
		GoldPerSecond            int           `json:"goldPerSecond"`
		JungleMinionsKilled      int           `json:"jungleMinionsKilled"`
		Level                    int           `json:"level"`
		MinionsKilled            int           `json:"minionsKilled"`
		ParticipantID            int           `json:"participantId"`
		Position                 Position      `json:"position"`
		TimeEnemySpentControlled int           `json:"timeEnemySpentControlled"`
		TotalGold                int           `json:"totalGold"`
		XP                       int           `json:"xp"`
	}

	ChampionStats struct {
		AbilityHaste         int `json:"abilityHaste"`
		AbilityPower         int `json:"abilityPower"`
		Armor                int `json:"armor"`
		ArmorPen             int `json:"armorPen"`
		ArmorPenPercent      int `json:"armorPenPercent"`
		AttackDamage         int `json:"attackDamage"`
		AttackSpeed          int `json:"attackSpeed"`
		BonusArmorPenPercent int `json:"bonusArmorPenPercent"`
		BonusMagicPenPercent int `json:"bonusMagicPenPercent"`
		CCReduction          int `json:"ccReduction"`
		CooldownReduction    int `json:"cooldownReduction"`
		Health               int `json:"health"`
		HealthMax            int `json:"healthMax"`
		HealthRegen          int `json:"healthRegen"`
		Lifesteal            int `json:"lifesteal"`
		MagicPen             int `json:"magicPen"`
		MagicPenPercent      int `json:"magicPenPercent"`
		MagicResist          int `json:"magicResist"`
		MovementSpeed        int `json:"movementSpeed"`
		Omnivamp             int `json:"omnivamp"`
		PhysicalVamp         int `json:"physicalVamp"`
		Power                int `json:"power"`
		PowerMax             int `json:"powerMax"`
		PowerRegen           int `json:"powerRegen"`
		SpellVamp            int `json:"spellVamp"`
	}

	DamageStats struct {
		MagicDamageDone               int `json:"magicDamageDone"`
		MagicDamageDoneToChampions    int `json:"magicDamageDoneToChampions"`
		MagicDamageTaken              int `json:"magicDamageTaken"`
		PhysicalDamageDone            int `json:"physicalDamageDone"`
		PhysicalDamageDoneToChampions int `json:"physicalDamageDoneToChampions"`
		PhysicalDamageTaken           int `json:"physicalDamageTaken"`
		TotalDamageDone               int `json:"totalDamageDone"`
		TotalDamageDoneToChampions    int `json:"totalDamageDoneToChampions"`
		TotalDamageTaken              int `json:"totalDamageTaken"`
		TrueDamageDone                int `json:"trueDamageDone"`
		TrueDamageDoneToChampions     int `json:"trueDamageDoneToChampions"`
		TrueDamageTaken               int `json:"trueDamageTaken"`
	}

	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
	}
)
