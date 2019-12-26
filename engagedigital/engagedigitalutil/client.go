package utils

import (
	"fmt"

	"github.com/grokify/oauth2more"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
)

const (
	ApiUrlFormat = `https://%s.api.engagement.dimelo.com/1.0`
)

func NewApiClient(site, token string) *engagedigital.APIClient {
	cfg := engagedigital.NewConfiguration()
	cfg.HTTPClient = oauth2more.NewClientBearerTokenSimple(token)
	cfg.BasePath = fmt.Sprintf(ApiUrlFormat, site)
	return engagedigital.NewAPIClient(cfg)
}
