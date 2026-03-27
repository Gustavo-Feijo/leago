package ddragon

import (
	"context"
	"fmt"
	"time"

	"github.com/Gustavo-Feijo/leago/api/ddragon/champion"
	"github.com/Gustavo-Feijo/leago/api/ddragon/item"
	"github.com/Gustavo-Feijo/leago/api/ddragon/languages"
	"github.com/Gustavo-Feijo/leago/api/ddragon/profileicon"
	"github.com/Gustavo-Feijo/leago/api/ddragon/realms"
	"github.com/Gustavo-Feijo/leago/api/ddragon/summoner"
	"github.com/Gustavo-Feijo/leago/api/ddragon/versions"
	"github.com/Gustavo-Feijo/leago/internal"
)

// RegionClient groups the DDragon APIs, also holding some params used on the requests (Like locale and version).
type RegionClient struct {
	version string
	locale  string

	Champion    *champion.RegionClient
	Item        *item.RegionClient
	Language    *languages.RegionClient
	ProfileIcon *profileicon.RegionClient
	Realms      *realms.RegionClient
	Summoner    *summoner.RegionClient
	Version     *versions.RegionClient
}

func NewRegionClient(baseClient *internal.Client, realm realms.Realm, version, locale string) (*RegionClient, error) {
	c := &RegionClient{
		Language: languages.NewRegionClient(baseClient),
		Realms:   realms.NewRegionClient(baseClient),
		Version:  versions.NewRegionClient(baseClient),
	}

	// If version or locale wasn't passed get the default through bootstrap.
	if version == "" || locale == "" {
		if err := c.bootstrap(realm); err != nil {
			return nil, fmt.Errorf("leago: DDragon bootstrap failed: %w", err)
		}
	}

	// Manual takes precedence.
	if version != "" {
		c.version = version
	}

	if locale != "" {
		c.locale = locale
	}

	c.Champion = champion.NewRegionClient(baseClient, c.version, c.locale)
	c.Item = item.NewRegionClient(baseClient, c.version, c.locale)
	c.ProfileIcon = profileicon.NewRegionClient(baseClient, c.version, c.locale)
	c.Summoner = summoner.NewRegionClient(baseClient, c.version, c.locale)

	return c, nil
}

// bootstrap gets the version and locale for the provided realm.
func (rc *RegionClient) bootstrap(realm realms.Realm) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	info, err := rc.Realms.GetRealm(ctx, realm)
	if err != nil {
		return err
	}

	rc.version = info.V
	rc.locale = info.L

	return nil
}
