package lite

import (
	"net/http"
	"strings"

	"github.com/grokify/gotilla/net/httputilmore"
)

const (
	EnvEngageVoiceServerURL = "ENGAGE_VOICE_SERVER_URL"
	EnvEngageVoiceAccountID = "ENGAGE_VOICE_ACCOUNT_ID"
	EnvEngageVoiceToken     = "ENGAGE_VOICE_TOKEN"
	APIUsersURL             = "/api/v1/admin/users"
	APITokensURL            = "/api/v1/admin/token"
)

type ClientLite struct {
	ServerURL  string
	Token      string
	AccountID  string
	HTTPClient *http.Client
}

func NewClientLite(serverURL, accountID, token string) ClientLite {
	client := ClientLite{
		ServerURL: strings.TrimSpace(serverURL),
		AccountID: strings.TrimSpace(accountID),
		Token:     strings.TrimSpace(token)}
	client.LoadHTTPClient(client.Token)
	return client
}

func (lite *ClientLite) LoadHTTPClient(token string) {
	if len(strings.TrimSpace(token)) == 0 {
		token = lite.Token
	}
	lite.HTTPClient = NewHTTPClient(token)
}

func (lite *ClientLite) Tokens() ([]string, error) {
	return ListTokens(lite.Token)
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
