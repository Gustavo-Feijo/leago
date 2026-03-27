package item

type (
	ItemResponse struct {
		Type    string          `json:"type"`
		Version string          `json:"version"`
		Basic   BasicItem       `json:"basic"`
		Data    map[string]Item `json:"data"`
		Groups  []ItemGroup     `json:"groups"`
		Tree    []ItemTree      `json:"tree"`
	}

	BasicItem struct {
		Name             string          `json:"name"`
		Rune             Rune            `json:"rune"`
		Gold             Gold            `json:"gold"`
		Group            string          `json:"group"`
		Description      string          `json:"description"`
		Colloq           string          `json:"colloq"`
		Plaintext        string          `json:"plaintext"`
		Consumed         bool            `json:"consumed"`
		Stacks           int             `json:"stacks"`
		Depth            int             `json:"depth"`
		ConsumeOnFull    bool            `json:"consumeOnFull"`
		From             []string        `json:"from"`
		Into             []string        `json:"into"`
		SpecialRecipe    int             `json:"specialRecipe"`
		InStore          bool            `json:"inStore"`
		HideFromAll      bool            `json:"hideFromAll"`
		RequiredChampion string          `json:"requiredChampion"`
		RequiredAlly     string          `json:"requiredAlly"`
		Stats            ItemStats       `json:"stats"`
		Tags             []string        `json:"tags"`
		Maps             map[string]bool `json:"maps"`
	}

	ItemStats struct {
		FlatHPPoolMod                       float64 `json:"FlatHPPoolMod"`
		RFlatHPModPerLevel                  float64 `json:"rFlatHPModPerLevel"`
		FlatMPPoolMod                       float64 `json:"FlatMPPoolMod"`
		RFlatMPModPerLevel                  float64 `json:"rFlatMPModPerLevel"`
		PercentHPPoolMod                    float64 `json:"PercentHPPoolMod"`
		PercentMPPoolMod                    float64 `json:"PercentMPPoolMod"`
		FlatHPRegenMod                      float64 `json:"FlatHPRegenMod"`
		RFlatHPRegenModPerLevel             float64 `json:"rFlatHPRegenModPerLevel"`
		PercentHPRegenMod                   float64 `json:"PercentHPRegenMod"`
		FlatMPRegenMod                      float64 `json:"FlatMPRegenMod"`
		RFlatMPRegenModPerLevel             float64 `json:"rFlatMPRegenModPerLevel"`
		PercentMPRegenMod                   float64 `json:"PercentMPRegenMod"`
		FlatArmorMod                        float64 `json:"FlatArmorMod"`
		RFlatArmorModPerLevel               float64 `json:"rFlatArmorModPerLevel"`
		PercentArmorMod                     float64 `json:"PercentArmorMod"`
		RFlatArmorPenetrationMod            float64 `json:"rFlatArmorPenetrationMod"`
		RFlatArmorPenetrationModPerLevel    float64 `json:"rFlatArmorPenetrationModPerLevel"`
		RPercentArmorPenetrationMod         float64 `json:"rPercentArmorPenetrationMod"`
		RPercentArmorPenetrationModPerLevel float64 `json:"rPercentArmorPenetrationModPerLevel"`
		FlatPhysicalDamageMod               float64 `json:"FlatPhysicalDamageMod"`
		RFlatPhysicalDamageModPerLevel      float64 `json:"rFlatPhysicalDamageModPerLevel"`
		PercentPhysicalDamageMod            float64 `json:"PercentPhysicalDamageMod"`
		FlatMagicDamageMod                  float64 `json:"FlatMagicDamageMod"`
		RFlatMagicDamageModPerLevel         float64 `json:"rFlatMagicDamageModPerLevel"`
		PercentMagicDamageMod               float64 `json:"PercentMagicDamageMod"`
		FlatMovementSpeedMod                float64 `json:"FlatMovementSpeedMod"`
		RFlatMovementSpeedModPerLevel       float64 `json:"rFlatMovementSpeedModPerLevel"`
		PercentMovementSpeedMod             float64 `json:"PercentMovementSpeedMod"`
		RPercentMovementSpeedModPerLevel    float64 `json:"rPercentMovementSpeedModPerLevel"`
		FlatAttackSpeedMod                  float64 `json:"FlatAttackSpeedMod"`
		PercentAttackSpeedMod               float64 `json:"PercentAttackSpeedMod"`
		RPercentAttackSpeedModPerLevel      float64 `json:"rPercentAttackSpeedModPerLevel"`
		RFlatDodgeMod                       float64 `json:"rFlatDodgeMod"`
		RFlatDodgeModPerLevel               float64 `json:"rFlatDodgeModPerLevel"`
		PercentDodgeMod                     float64 `json:"PercentDodgeMod"`
		FlatCritChanceMod                   float64 `json:"FlatCritChanceMod"`
		RFlatCritChanceModPerLevel          float64 `json:"rFlatCritChanceModPerLevel"`
		PercentCritChanceMod                float64 `json:"PercentCritChanceMod"`
		FlatCritDamageMod                   float64 `json:"FlatCritDamageMod"`
		RFlatCritDamageModPerLevel          float64 `json:"rFlatCritDamageModPerLevel"`
		PercentCritDamageMod                float64 `json:"PercentCritDamageMod"`
		FlatBlockMod                        float64 `json:"FlatBlockMod"`
		PercentBlockMod                     float64 `json:"PercentBlockMod"`
		FlatSpellBlockMod                   float64 `json:"FlatSpellBlockMod"`
		RFlatSpellBlockModPerLevel          float64 `json:"rFlatSpellBlockModPerLevel"`
		PercentSpellBlockMod                float64 `json:"PercentSpellBlockMod"`
		FlatEXPBonus                        float64 `json:"FlatEXPBonus"`
		PercentEXPBonus                     float64 `json:"PercentEXPBonus"`
		RPercentCooldownMod                 float64 `json:"rPercentCooldownMod"`
		RPercentCooldownModPerLevel         float64 `json:"rPercentCooldownModPerLevel"`
		RFlatTimeDeadMod                    float64 `json:"rFlatTimeDeadMod"`
		RFlatTimeDeadModPerLevel            float64 `json:"rFlatTimeDeadModPerLevel"`
		RPercentTimeDeadMod                 float64 `json:"rPercentTimeDeadMod"`
		RPercentTimeDeadModPerLevel         float64 `json:"rPercentTimeDeadModPerLevel"`
		RFlatGoldPer10Mod                   float64 `json:"rFlatGoldPer10Mod"`
		RFlatMagicPenetrationMod            float64 `json:"rFlatMagicPenetrationMod"`
		RFlatMagicPenetrationModPerLevel    float64 `json:"rFlatMagicPenetrationModPerLevel"`
		RPercentMagicPenetrationMod         float64 `json:"rPercentMagicPenetrationMod"`
		RPercentMagicPenetrationModPerLevel float64 `json:"rPercentMagicPenetrationModPerLevel"`
		FlatEnergyRegenMod                  float64 `json:"FlatEnergyRegenMod"`
		RFlatEnergyRegenModPerLevel         float64 `json:"rFlatEnergyRegenModPerLevel"`
		FlatEnergyPoolMod                   float64 `json:"FlatEnergyPoolMod"`
		RFlatEnergyModPerLevel              float64 `json:"rFlatEnergyModPerLevel"`
		PercentLifeStealMod                 float64 `json:"PercentLifeStealMod"`
		PercentSpellVampMod                 float64 `json:"PercentSpellVampMod"`
	}

	Item struct {
		Name        string             `json:"name"`
		Description string             `json:"description"`
		Colloq      string             `json:"colloq"`
		Plaintext   string             `json:"plaintext"`
		Into        []string           `json:"into,omitempty"`
		From        []string           `json:"from,omitempty"`
		Image       Image              `json:"image"`
		Gold        Gold               `json:"gold"`
		Tags        []string           `json:"tags"`
		Maps        map[string]bool    `json:"maps"`
		Stats       map[string]float64 `json:"stats"` // dynamic stats (varies per item)
	}

	Rune struct {
		IsRune bool   `json:"isrune"`
		Tier   int    `json:"tier"`
		Type   string `json:"type"`
	}

	Gold struct {
		Base        int  `json:"base"`
		Total       int  `json:"total"`
		Sell        int  `json:"sell"`
		Purchasable bool `json:"purchasable"`
	}

	Image struct {
		Full   string `json:"full"`
		Sprite string `json:"sprite"`
		Group  string `json:"group"`
		X      int    `json:"x"`
		Y      int    `json:"y"`
		W      int    `json:"w"`
		H      int    `json:"h"`
	}

	ItemGroup struct {
		ID              string `json:"id"`
		MaxGroupOwnable string `json:"MaxGroupOwnable"`
	}

	ItemTree struct {
		Header string   `json:"header"`
		Tags   []string `json:"tags"`
	}
)
