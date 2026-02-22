package championmastery

import (
	"context"
	"fmt"
	"leago/internal"
	"net/http"
	"net/url"
)

type (
	Player struct {
		Puuid    string `json:"puuid"`
		TeamID   string `json:"teamId"`
		Position string `json:"position"` // UNSELECTED, FILL, TOP, JUNGLE, MIDDLE, BOTTOM, UTILITY
		Role     string `json:"role"`     // CAPTAIN, MEMBER
	}

	PlayersResponse []Player
)

// GetByPUUID returns the player champion mastery informations got by their PUUID.
func (pc *PlatformClient) GetByPUUID(ctx context.Context, puuid string) (PlayersResponse, error) {
	endpoint := fmt.Sprintf(
		"/lol/clash/v1/players/by-puuid/%s",
		url.PathEscape(puuid),
	)

	uri := pc.client.GetURL(endpoint)
	return internal.AuthRequest[PlayersResponse](ctx, pc.client, pc.client.ApiKey, uri, http.MethodGet, nil, internal.WithMethod("Clash.GetByPUUID"))
}
