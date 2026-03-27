package champion

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gustavo-Feijo/leago/internal"
	"github.com/Gustavo-Feijo/leago/options"
)

const (
	MethodGetChampions    = "DDChampion.GetChampions"
	MethodGetChampionByID = "DDChampion.GetChampionByID"
)

// GetChampions returns a list of champions and their informations.
func (rc *RegionClient) GetChampions(
	ctx context.Context,
	opts ...options.PublicOption,
) (ChampionResponse, error) {
	endpoint := fmt.Sprintf("/cdn/%s/data/%s/champion.json", rc.version, rc.language)

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetChampions),
	}

	uri := rc.client.GetDDragonURL(endpoint)
	return internal.Request[ChampionResponse](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}

// GetChampionByID returns the information for a given champion. ID is a string like "Aatrox", not the numerical key.
func (rc *RegionClient) GetChampionByID(
	ctx context.Context,
	championID string,
	opts ...options.PublicOption,
) (SingleChampionResponse, error) {
	endpoint := fmt.Sprintf("/cdn/%s/data/%s/champion/%s.json", rc.version, rc.language, url.PathEscape(championID))

	defaultOpts := []internal.RequestOption{
		internal.WithAPIMethod(MethodGetChampions),
	}

	uri := rc.client.GetDDragonURL(endpoint)
	return internal.Request[SingleChampionResponse](
		ctx,
		rc.client,
		uri,
		options.MergeOptions(defaultOpts, opts)...,
	)
}
