package utils

import (
	"fmt"

	"github.com/grokify/oauth2more"

	engagedigital "github.com/grokify/go-ringcentral-client/engagedigital/v1/client"
)

const (
	ApiUrlFormat = `https://%s.api.engagement.dimelo.com/1.0`
)

func NewApiClient(site, token string) *engagedigital.APIClient {
	cfg := engagedigital.NewConfiguration()
	cfg.HTTPClient = oauth2more.NewClientToken(oauth2more.TokenBearer, token, false)
	cfg.BasePath = fmt.Sprintf(ApiUrlFormat, site)
	return engagedigital.NewAPIClient(cfg)
}
