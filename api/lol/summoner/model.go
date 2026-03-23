package summoner

import "github.com/Gustavo-Feijo/leago/internal"

// Summoner holds the public profile data for a League of Legends account.
type Summoner struct {
	ProfileIconID int                     `json:"profileIconId"`
	RevisionDate  internal.UnixMillisTime `json:"revisionDate"`
	Puuid         string                  `json:"puuid"`
	SummonerLevel int64                   `json:"summonerLevel"`
}
