package champion

type (
	ChampionResponse struct {
		Type    string                  `json:"type"`
		Format  string                  `json:"format"`
		Version string                  `json:"version"`
		Data    map[string]ChampionData `json:"data"`
	}

	ChampionData struct {
		Version string   `json:"version"`
		ID      string   `json:"id"`
		Key     string   `json:"key"`
		Name    string   `json:"name"`
		Title   string   `json:"title"`
		Blurb   string   `json:"blurb"`
		Info    Info     `json:"info"`
		Image   Image    `json:"image"`
		Tags    []string `json:"tags"`
		Partype string   `json:"partype"`
		Stats   Stats    `json:"stats"`
	}
)

type (
	SingleChampionResponse struct {
		Type    string                      `json:"type"`
		Format  string                      `json:"format"`
		Version string                      `json:"version"`
		Data    map[string]FullChampionData `json:"data"`
	}

	FullChampionData struct {
		ID          string   `json:"id"`
		Key         string   `json:"key"`
		Name        string   `json:"name"`
		Title       string   `json:"title"`
		Image       Image    `json:"image"`
		Skins       []Skin   `json:"skins"`
		Lore        string   `json:"lore"`
		Blurb       string   `json:"blurb"`
		AllyTips    []string `json:"allytips"`
		EnemyTips   []string `json:"enemytips"`
		Tags        []string `json:"tags"`
		Partype     string   `json:"partype"`
		Info        Info     `json:"info"`
		Stats       Stats    `json:"stats"`
		Spells      []Spell  `json:"spells"`
		Passive     Passive  `json:"passive"`
		Recommended []any    `json:"recommended"`
	}

	Skin struct {
		ID         string `json:"id"`
		Num        int    `json:"num"`
		Name       string `json:"name"`
		Chromas    bool   `json:"chromas"`
		ParentSkin int    `json:"parentSkin,omitempty"`
	}

	Spell struct {
		ID           string         `json:"id"`
		Name         string         `json:"name"`
		Description  string         `json:"description"`
		Tooltip      string         `json:"tooltip"`
		LevelTip     LevelTip       `json:"leveltip"`
		MaxRank      int            `json:"maxrank"`
		Cooldown     []float64      `json:"cooldown"`
		CooldownBurn string         `json:"cooldownBurn"`
		Cost         []float64      `json:"cost"`
		CostBurn     string         `json:"costBurn"`
		DataValues   map[string]any `json:"datavalues"`
		Effect       [][]float64    `json:"effect"`
		EffectBurn   []string       `json:"effectBurn"`
		Vars         []any          `json:"vars"`
		CostType     string         `json:"costType"`
		MaxAmmo      string         `json:"maxammo"`
		Range        []float64      `json:"range"`
		RangeBurn    string         `json:"rangeBurn"`
		Image        Image          `json:"image"`
		Resource     string         `json:"resource"`
	}

	LevelTip struct {
		Label  []string `json:"label"`
		Effect []string `json:"effect"`
	}

	Passive struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Image       Image  `json:"image"`
	}
)

type (
	Info struct {
		Attack     int `json:"attack"`
		Defense    int `json:"defense"`
		Magic      int `json:"magic"`
		Difficulty int `json:"difficulty"`
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

	Stats struct {
		HP                   float64 `json:"hp"`
		HPPerLevel           float64 `json:"hpperlevel"`
		MP                   float64 `json:"mp"`
		MPPerLevel           float64 `json:"mpperlevel"`
		MoveSpeed            float64 `json:"movespeed"`
		Armor                float64 `json:"armor"`
		ArmorPerLevel        float64 `json:"armorperlevel"`
		SpellBlock           float64 `json:"spellblock"`
		SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
		AttackRange          float64 `json:"attackrange"`
		HPRegen              float64 `json:"hpregen"`
		HPRegenPerLevel      float64 `json:"hpregenperlevel"`
		MPRegen              float64 `json:"mpregen"`
		MPRegenPerLevel      float64 `json:"mpregenperlevel"`
		Crit                 float64 `json:"crit"`
		CritPerLevel         float64 `json:"critperlevel"`
		AttackDamage         float64 `json:"attackdamage"`
		AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
		AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
		AttackSpeed          float64 `json:"attackspeed"`
	}
)
