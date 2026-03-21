package content

type (
	Content struct {
		Version      string        `json:"version"`
		Characters   []ContentItem `json:"characters"`
		Maps         []ContentItem `json:"maps"`
		Chromas      []ContentItem `json:"chromas"`
		Skins        []ContentItem `json:"skins"`
		SkinLevels   []ContentItem `json:"skinLevels"`
		Equips       []ContentItem `json:"equips"`
		GameModes    []ContentItem `json:"gameModes"`
		Sprays       []ContentItem `json:"sprays"`
		SprayLevels  []ContentItem `json:"sprayLevels"`
		Charms       []ContentItem `json:"charms"`
		CharmLevels  []ContentItem `json:"charmLevels"`
		PlayerCards  []ContentItem `json:"playerCards"`
		PlayerTitles []ContentItem `json:"playerTitles"`
		Acts         []Act         `json:"acts"`
	}

	ContentItem struct {
		Name           string          `json:"name"`
		LocalizedNames *LocalizedNames `json:"localizedNames,omitempty"`
		ID             string          `json:"id"`
		AssetName      string          `json:"assetName"`
		AssetPath      *string         `json:"assetPath,omitempty"`
	}

	LocalizedNames struct {
		ArAE string `json:"ar-AE"`
		DeDE string `json:"de-DE"`
		EnGB string `json:"en-GB"`
		EnUS string `json:"en-US"`
		EsES string `json:"es-ES"`
		EsMX string `json:"es-MX"`
		FrFR string `json:"fr-FR"`
		IdID string `json:"id-ID"`
		ItIT string `json:"it-IT"`
		JaJP string `json:"ja-JP"`
		KoKR string `json:"ko-KR"`
		PlPL string `json:"pl-PL"`
		PtBR string `json:"pt-BR"`
		RuRU string `json:"ru-RU"`
		ThTH string `json:"th-TH"`
		TrTR string `json:"tr-TR"`
		ViVN string `json:"vi-VN"`
		ZhCN string `json:"zh-CN"`
		ZhTW string `json:"zh-TW"`
	}

	Act struct {
		Name           string          `json:"name"`
		LocalizedNames *LocalizedNames `json:"localizedNames,omitempty"`
		ID             string          `json:"id"`
		IsActive       bool            `json:"isActive"`
	}
)
