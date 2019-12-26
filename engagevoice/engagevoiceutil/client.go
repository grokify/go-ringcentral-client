package engagevoiceutil

import (
	"net/http"

	"github.com/grokify/go-ringcentral-engage/engagevoice/engagevoice"
	"github.com/grokify/gotilla/net/httputilmore"
)

type ClientAPIs struct {
	APIClient *engagevoice.APIClient
	Config    *engagevoice.Configuration
}

func NewClientAPIsHTTPClient(httpClient *http.Client) ClientAPIs {
	cfg := engagevoice.NewConfiguration()
	cfg.HTTPClient = httpClient
	return ClientAPIs{
		Config:    cfg,
		APIClient: engagevoice.NewAPIClient(cfg)}
}

func NewClientAPIs(apiToken string) ClientAPIs {
	httpClient := NewHTTPClient(apiToken)
	return NewClientAPIsHTTPClient(httpClient)
}

func NewHTTPClient(token string) *http.Client {
	header := http.Header{}
	header.Add(EngageVoiceTokenHeader, token)

	client := &http.Client{}
	client.Transport = httputilmore.TransportWithHeaders{
		Transport: client.Transport,
		Header:    header}
	return client
}
