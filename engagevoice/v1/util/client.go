package engagevoiceutil

import (
	"net/http"

	engagevoice "github.com/grokify/go-ringcentral-client/engagevoice/v1/client"
	"github.com/grokify/go-ringcentral-client/engagevoice/v1/util/lite"
)

type ClientAPIs struct {
	APIClient *engagevoice.APIClient
	Config    *engagevoice.Configuration
}

func NewClientAPIsRingCentralPassword(rcCredentials []byte) (ClientAPIs, error) {
	httpClient, err := lite.NewClientRingCentralPassword(rcCredentials)
	if err != nil {
		return ClientAPIs{}, err
	}
	return NewClientAPIsHTTPClient(httpClient), nil
}

func NewClientAPIsHTTPClient(httpClient *http.Client) ClientAPIs {
	cfg := engagevoice.NewConfiguration()
	cfg.HTTPClient = httpClient
	return ClientAPIs{
		Config:    cfg,
		APIClient: engagevoice.NewAPIClient(cfg)}
}

func NewClientAPIs(apiToken string) ClientAPIs {
	return NewClientAPIsHTTPClient(lite.NewClientToken(apiToken))
}
