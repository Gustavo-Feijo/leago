package realms

type (
	DDragonResponse struct {
		N              DDragonVersions `json:"n"`
		V              string          `json:"v"`
		L              string          `json:"l"`
		CDN            string          `json:"cdn"`
		DD             string          `json:"dd"`
		LG             string          `json:"lg"`
		CSS            string          `json:"css"`
		ProfileIconMax int             `json:"profileiconmax"`
		Store          *string         `json:"store"`
	}

	DDragonVersions struct {
		Item        string `json:"item"`
		Rune        string `json:"rune"`
		Mastery     string `json:"mastery"`
		Summoner    string `json:"summoner"`
		Champion    string `json:"champion"`
		ProfileIcon string `json:"profileicon"`
		Map         string `json:"map"`
		Language    string `json:"language"`
		Sticker     string `json:"sticker"`
	}
)
