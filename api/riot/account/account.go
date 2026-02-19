package account

import (
	"context"
	"fmt"
	"leago/internal"
	"net/http"
	"net/url"
)

type Account struct {
	Puuid    string `json:"puuid"`
	GameName string `json:"gameName"`
	TagLine  string `json:"tagLine"`
}

// GetByRiotID returns the user account by their gamename and tagline.
func (rc *RegionClient) GetByRiotID(ctx context.Context, gameName, tagLine string) (*Account, error) {
	endpoint := fmt.Sprintf(
		"/riot/account/v1/accounts/by-riot-id/%s/%s",
		url.PathEscape(gameName),
		url.PathEscape(tagLine),
	)

	uri := rc.client.GetURL(endpoint)
	return internal.AuthRequest[*Account](ctx, rc.client.Http, rc.client.ApiKey, uri, http.MethodGet, nil)
}
