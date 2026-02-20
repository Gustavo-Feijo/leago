package championmastery

import (
	"context"
	"fmt"
	"leago/internal"
	"net/http"
	"net/url"
	"strconv"
)

type (
	MasteryList []Mastery

	Mastery struct {
		Puuid                        string               `json:"puuid"`
		ChampionPointsUntilNextLevel int64                `json:"championPointsUntilNextLevel"`
		ChestGranted                 bool                 `json:"chestGranted"`
		ChampionID                   int64                `json:"championId"`
		LastPlayTime                 int64                `json:"lastPlayTime"`
		ChampionLevel                int                  `json:"championLevel"`
		ChampionPoints               int                  `json:"championPoints"`
		ChampionPointsSinceLastLevel int64                `json:"championPointsSinceLastLevel"`
		MarkRequiredForNextLevel     int                  `json:"markRequiredForNextLevel"`
		ChampionSeasonMilestone      int                  `json:"championSeasonMilestone"`
		NextSeasonMilestone          NextSeasonMilestones `json:"nextSeasonMilestone"`
		TokensEarned                 int                  `json:"tokensEarned"`
		MilestoneGrades              []string             `json:"milestoneGrades"`
	}

	NextSeasonMilestones struct {
		RequireGradeCounts map[string]int `json:"requireGradeCounts"`
		RewardMarks        int            `json:"rewardMarks"`
		Bonus              bool           `json:"bonus"`
		RewardConfig       RewardConfig   `json:"rewardConfig"`
	}

	RewardConfig struct {
		RewardValue   string `json:"rewardValue"`
		RewardType    string `json:"rewardType"`
		MaximumReward int    `json:"maximumReward"`
	}

	MasteryScore int
)

// GetByPUUID returns the player champion mastery informations got by their PUUID.
func (pc *PlatformClient) GetByPUUID(ctx context.Context, puuid string) (MasteryList, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/champion-masteries/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[MasteryList](ctx, pc.client.Http, pc.client.ApiKey, uri, http.MethodGet, nil)
}

// GetByPUUIDTop returns the top X player champion mastery informations got by their PUUID.
func (pc *PlatformClient) GetByPUUIDTop(ctx context.Context, puuid string, count int) (MasteryList, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/top",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)

	params := make(map[string]string)
	params["count"] = strconv.Itoa(count)

	return internal.AuthRequest[MasteryList](ctx, pc.client.Http, pc.client.ApiKey, uri, http.MethodGet, params)
}

// GetByPUUIDByChampion returns the player champion mastery informations for a given champion got by their PUUID.
func (pc *PlatformClient) GetByPUUIDByChampion(ctx context.Context, championID int64, puuid string) (Mastery, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/by-champion/%d",
		url.PathEscape(puuid),
		championID,
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[Mastery](ctx, pc.client.Http, pc.client.ApiKey, uri, http.MethodGet, nil)
}

// GetScoreByPUUID returns a player total champion mastery score (Sum of individual champion mastery levels).
func (pc *PlatformClient) GetScoreByPUUID(ctx context.Context, puuid string) (MasteryScore, error) {
	endpoint := fmt.Sprintf(
		"/lol/champion-mastery/v4/scores/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[MasteryScore](ctx, pc.client.Http, pc.client.ApiKey, uri, http.MethodGet, nil)
}
